import { Resource, component$ } from "@builder.io/qwik";
import { useEndpoint } from "@builder.io/qwik-city";
import type { RequestHandler } from "@builder.io/qwik-city";
import axios from "axios";

export interface ModuleSummary {
  Name: string
  Providers: {
    Source: string
  }[]
}

export const onGet: RequestHandler<ModuleSummary[]> = async () => {
  let modules: Array<ModuleSummary> = []
  try {
    const response = await axios.get(`http://back:8080/modules`);
    modules = await response.data;
  } catch (error) {
    console.log(error);
  }
  return modules;
};


export default component$(() => {
  const moduleList = useEndpoint<ModuleSummary[]>();

  return (
    <Resource
      value={moduleList}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
      onResolved={(moduleList) => (
        <div class="block">
          {moduleList.map(module => {
            return (
              <div class="block container">
                <a href={`/modules/${module.Name}`} class="box">
                  <div>
                    <div class="title is-5">{module.Name}</div>
                    {module.Providers.map(prov => {
                      return (
                        <div class="field is-inline m-1 tags">
                          <span class="tag is-info">
                            {prov.Source}
                          </span>
                        </div>
                      )
                    })}
                  </div>
                </a>
              </div>
            )
          })}
        </div>
      )}
    />
  );
});