import { Resource, component$ } from "@builder.io/qwik";
import { useEndpoint } from "@builder.io/qwik-city";
import type { RequestHandler } from "@builder.io/qwik-city";
import axios from "axios";

export interface ModuleSource {
    ID: string
    name: string
    type: string
    address: string
    group: string
    auth: string
}

export const onGet: RequestHandler<ModuleSource[]> = async () => {
  let sources: Array<ModuleSource> = []
  try {
    const response = await axios.get(`http://back:8080/config`);
    sources = await response.data;
    console.log(sources)
  } catch (error) {
    console.log(error);
  }
  return sources;
};


export default component$(() => {
  const sourceList = useEndpoint<ModuleSource[]>();

  return (
    <Resource
      value={sourceList}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
        onResolved={(sourceList) => (
        <div class="block content-center mx-auto max-w-xl">
          {sourceList.map(source => {
            return (
              <a href={`/settings/${source.name}`} class="box">
                <div>
                  <div class="title is-5">{source.name}</div>
                    <span class="tag is-info">
                      {source.type}
                    </span>
                </div>
              </a>
            )
          })}
        </div>
      )}
    />
  );
});