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
    <div class="tab-pane">
      {variables.map(variable => {
        return (
          <details class="card m-1">
            <summary class="">
              <div class="m-1">
                <h4 class="title is-5">{variable.Name}</h4>
                <span class="tag">
                  {variable.Type}
                </span>
                {variable.Required ?
                  <span class="tag is-danger">
                    required
                  </span>
                  : <></>}
              </div>
            </summary>
            <dl class="card-content">
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
    <div class="tab-pane">
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
  const active: string = "is-active";
  const inactive: string = "";
  return (
    <Resource
      value={moduleData}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
      onResolved={(module) => (
        <div class="box">
          <div class="">
            <h2 class="title is-2">{module.Name}</h2>
            <p class="title is-6">ID: {module.ID}</p>
            <div class="tags has-addons">
              {module.Providers.map(prov => {
                return (
                  <div class="m-1">
                    <span class="tag">{prov.Source}</span>
                    <span class="tag is-primary">
                      {prov.VersionConstraints}
                    </span>
                  </div>
                );
              })}
            </div>
          </div>
          <div class="mt-3 m-5">
            <div class="tabs is-toggle is-centered main-menu is-medium">
              <ul>
                <li class={store.tab === "variables" ? active : inactive}>
                  <a>
                  <span onClick$={() => store.tab = "variables"}>
                    Variables
                  </span>
                  </a>
                </li>
                <li class={store.tab === "outputs" ? active : inactive}>
                  <a>
                  <span onClick$={() => store.tab = "outputs"}>
                    Outputs
                  </span>
                  </a>
                </li>
              </ul>
            </div>
            <div class="tab-content">
              {store.tab === "variables" ? variablesTab(module.Variables) : <></>}
              {store.tab === "outputs" ? outputsTab(module.Outputs) : <></>}
            </div>
          </div>
        </div>
      )}
    />
  );
});
