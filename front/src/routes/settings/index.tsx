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
    const response = await axios.get(`http://localhost:8080/config`);
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
              <a href={`/settings/${source.name}`} class="m-2 group block mx-auto rounded-sm p-4 bg-white ring-1 hover:bg-blue-600 hover:ring-blue-600">
                <div>
                  <div class="group-hover:text-white text-lg font-medium text-black">{source.name}</div>
                    <span class="text-xs font-regular inline-block py-1 px-2 rounded lowercase last:mr-0 mr-1  group-hover:text-blue-500 bg-blue-100 group-hover:font-semibold">
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