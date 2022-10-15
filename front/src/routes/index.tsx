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
    const response = await axios.get(`http://localhost:8080/modules`);
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
        <div class="block content-center mx-auto max-w-xl">
          {moduleList.map(module => {
            return (
              <a href={`/modules/${module.Name}`} class="m-2 group block mx-auto rounded-sm p-4 bg-white ring-1 hover:bg-indigo-600 hover:ring-indigo-600">
                <div>
                  <div class="group-hover:text-white text-lg font-medium text-black">{module.Name}</div>
                    {module.Providers.map(prov => {
                      return(
                        <span class="text-xs font-semibold inline-block py-1 px-2 rounded text-indigo-500 bg-indigo-200 lowercase last:mr-0 mr-1">
                          {prov.Source}
                        </span>
                      )
                    })}
                </div>
              </a>
            )
          })}
        </div>
      )}
    />
  );
});