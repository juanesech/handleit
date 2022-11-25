import { Resource, component$, useStore } from "@builder.io/qwik";
import { useEndpoint } from "@builder.io/qwik-city";
import type { RequestHandler } from "@builder.io/qwik-city";
import axios from "axios";

interface Module {
  ID: string
  Name: string
  Variables: Variable[]
  Outputs: Output[]
  Providers: {
    Source: string
    VersionConstraints: string[]
  }[]
}

interface Variable {
  Name: string
  Type: string
  Description: string
  Default: string
  Required: boolean
}

interface Output {
  Name: string
  Description: string
}

export const onGet: RequestHandler<Module> = async ({ params }) => {
  try {
    const response = await axios.get(`http://back:8080/modules/${params.module}`);
    const data = await response.data;
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

export const variablesTab = (variables: Variable[]) => {
  return (
    <div class="rounded-sm">
      {variables.map(variable => {
        return (
          <details>
            <summary class="bg-gray-100 group cursor-pointer hover:bg-blue-600 hover:ring-blue-600 list-none flex flex-wrap items-center rounded-sm">
              <div class="rounded-sm px-4 py-3 sm:px-6 group-hover:bg-blue-600 hover:ring-blue-600">
                <h4 class="group-hover:text-white text-md text-gray-600">{variable.Name}</h4>
                <span class="align-sub text-xs font-medium inline-block py-1 px-2 rounded group-hover:text-semibold group-hover:text-blue-500 bg-blue-100 lowercase last:mr-0 mr-1">
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
        );
      })}
    </div>
  );
}

export const outputsTab = (outputs: Output[]) => {
  return (
    <div class="rounded-sm">
      {outputs.map(output => {
        return (
            <dl class="rounded-sm">
              <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                <dt class="text-md font-medium text-gray-500">Name</dt>
                <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{output.Name}</dd>
              </div>
              <div class=" px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                <dt class="text-md font-medium text-gray-500">Description</dt>
                <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{output.Description}</dd>
              </div>
            </dl>
        );
      })}
    </div>
  );
}

export default component$(() => {
  const moduleData = useEndpoint<Module>();
  const store = useStore({ tab: "variables" });
  const active:string = "text-white px-4 py-5 sm:px-6 text-white bg-blue-600 ring-blue-600 rounded-sm";
  const inactive:string = "px-4 py-5 sm:px-6 cursor-pointer rounded-sm text-blue-600";
  return (
    <Resource
      value={moduleData}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
      onResolved={(module) => (
        <div class="overflow-hidden bg-white rounded-sm -mt-">
          <div class="px-4 py-3 sm:px-6">
            <h2 class="text-5xl font-light leading-6 text-blue-700 py-4">{module.Name}</h2>
            <p class="mt-1 max-w-2xl text-md text-gray-500">{module.ID}</p>
            {module.Providers.map(prov => {
              return (
                <span class="text-xs font-semibold inline-block py-1 px-2 rounded text-blue-500 bg-blue-200 lowercase last:mr-0 mr-1">
                  {`${prov.Source} ${prov.VersionConstraints}`}
                </span>
              );
            })}
          </div>
          <div>
            <ul class="flex flex-wrap text-sm font-medium text-center py-0">
              <li class="mr-2">
                <div class={ store.tab === "variables"? active: inactive}
                  onClick$={ () => store.tab = "variables"}>
                  <h2 class="text-2xl font-light leading-3">Variables</h2>
                </div>
              </li>
              <li class="mr-2">
                <div class={ store.tab === "outputs"? active: inactive}
                  onClick$={ () => store.tab = "outputs"}>
                  <h2 class="text-2xl font-light leading-3">Outputs</h2>
                </div>
              </li>
            </ul>
          </div>
          { store.tab === "variables"? variablesTab(module.Variables): <></> }
          { store.tab === "outputs"? outputsTab(module.Outputs): <></> }
        </div>
      )}
    />
  );
});
