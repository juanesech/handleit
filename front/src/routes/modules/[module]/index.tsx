import { Resource, component$, useStore } from "@builder.io/qwik";
import { useEndpoint } from "@builder.io/qwik-city";
import type { RequestHandler } from "@builder.io/qwik-city";
import axios from "axios";

interface Module {
  ID: string;
  Name: string;
  Variables: {
    Name: string
    Type: string
    Description: string
    Default: string
    Required: boolean
  }[],
  Outputs: {
    Name: string
    Description: string
  }[],
  Providers: {
    Source: string
    VersionConstraints: string[]
  }[]
}

export const onGet: RequestHandler<Module> = async ({ params }) => {
  try {
    const response = await axios.get(`http://localhost:8080/modules/${params.module}`);
    let data = await response.data;
    return {
      Name: data.Name,
      ID: data.ID,
      Variables: data.Variables,
      Outputs: data.Outputs,
      Providers: data.Providers
    };
  } catch (error) {
    console.log(error);
  }
};

export default component$(() => {
  const moduleData = useEndpoint<Module>();
  const store = useStore({ tab:"variables" })

  return (
    <Resource
      value={moduleData}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
      onResolved={(module) => (
        <div class="overflow-hidden bg-white sm:rounded-sm rounded-sm">
          <div class="px-4 py-5 sm:px-6">
            <h2 class="text-4xl font-light leading-6 text-indigo-700">{module.Name}</h2>
            <p class="mt-1 max-w-2xl text-md text-gray-500">{module.ID}</p>
            {module.Providers.map(prov => {
              return (
                <span class="text-xs font-semibold inline-block py-1 px-2 rounded text-indigo-500 bg-indigo-200 lowercase last:mr-0 mr-1">
                  {prov.Source}
                </span>
              )
            })}
          </div>
          <div>
            <ul class="flex flex-wrap text-sm font-medium text-center py-0">
              <li class="mr-2">
                <div class="px-4 py-5 sm:px-6 cursor-pointer border-t border-r border-l border-indigo-600 border-b-gray-50 border-b rounded-sm">
                  <h2 class="text-2xl font-light leading-3 text-indigo-600">Variables</h2>
                </div>
              </li>
              <li class="mr-2">
                <div class="px-4 py-5 sm:px-6 cursor-pointer border-t border-r border-l border-indigo-600 rounded-sm">
                  <h2 class="text-2xl font-light leading-3 text-indigo-600">Outputs</h2>
                </div>
              </li>
            </ul>
          </div>
          <div class="border border-indigo-600 rounded-sm">
            {module.Variables.map(variable => {
              return (
                <details>
                  <summary class="bg-gray-50 group rounded-sm cursor-pointer hover:bg-indigo-600 hover:ring-indigo-600 list-none flex flex-wrap items-center">
                    <div class="rounded-sm px-4 py-3 sm:px-6 group-hover:bg-indigo-600 hover:ring-indigo-600">
                      <h4 class="group-hover:text-white text-md text-gray-600">{variable.Name}</h4>
                      <span class="align-sub text-xs font-medium inline-block py-1 px-2 rounded text-indigo-500 bg-indigo-200 lowercase last:mr-0 mr-1">
                        {variable.Type}
                      </span>
                      {variable.Required ?
                        <span class="align-sub text-xs font-medium inline-block py-1 px-2 rounded text-red-500 bg-red-200 lowercase last:mr-0 mr-1">
                          required
                        </span>
                        : <></>}
                    </div>
                  </summary>
                  <dl>
                    <div class=" bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt class="text-xs font-medium text-gray-500">Description</dt>
                      <dd class="mt-1 text-xs text-gray-900 sm:col-span-2 sm:mt-0">{variable.Description}</dd>
                    </div>
                    <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt class="text-xs font-medium text-gray-500">Type</dt>
                      <dd class="mt-1 text-xs text-gray-900 sm:col-span-2 sm:mt-0">{variable.Type}</dd>
                    </div>
                    <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt class="text-xs font-medium text-gray-500">Required</dt>
                      <dd class="mt-1 text-xs text-gray-900 sm:col-span-2 sm:mt-0">{String(variable.Required)}</dd>
                    </div>
                    <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt class="text-xs font-medium text-gray-500">Default value</dt>
                      <dd class="mt-1 text-xs text-gray-900 sm:col-span-2 sm:mt-0">{variable.Default}</dd>
                    </div>
                  </dl>
                </details>
              )
            })}
          </div>
        </div>
      )}
    />
  );
});
