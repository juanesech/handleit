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
        <div class="block content-center mx-auto max-w-xl">
          {moduleList.map(module => {
            return (
              <a href={`/modules/${module.Name}`} class="m-2 group block mx-auto rounded-sm p-4 bg-white ring-1 hover:bg-blue-600 hover:ring-blue-600">
                <div>
                  <div class="group-hover:text-white text-lg font-medium text-black">{module.Name}</div>
                    {module.Providers.map(prov => {
                      return(
                        <span class="text-xs font-regular inline-block py-1 px-2 rounded lowercase last:mr-0 mr-1  group-hover:text-blue-500 bg-blue-100 group-hover:font-semibold">
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