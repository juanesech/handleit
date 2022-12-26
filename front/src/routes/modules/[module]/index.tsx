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
    <div>
      {variables.map(variable => {
        return (
          <div>
            <details class="box mb-2 is-list is-block">
              <summary class="is-clickable">
                <div class="m-1">
                  <h4 class="title is-5">{variable.Name}</h4>
                  <span class="tag mx-2">
                    {variable.Type}
                  </span>
                  {variable.Required ?
                    <span class="tag is-danger">
                      required
                    </span>
                    : <></>}
                </div>
              </summary>
              <table class="table is-striped is-hoverable is-relative">
                <tbody>
                  <tr>
                    <td class="is-small">Description</td>
                    <td class="">{variable.Description}</td>
                  </tr>
                  <tr>
                    <td class="text-xs font-medium text-gray-500">Type</td>
                    <td class=""><code>{variable.Type}</code></td>
                  </tr>
                  <tr>
                    <td class="text-xs font-medium text-gray-500">Required</td>
                    <td class="">{String(variable.Required)}</td></tr>
                  <tr>
                    <td class="text-xs font-medium text-gray-500">Default value</td>
                    <td class=""><code>{variable.Default}</code></td></tr>
                </tbody>
              </table>
            </details>
          </div>
        );
      })}
    </div>
  );
}

export const outputsTab = (outputs: Output[]) => {
  return (
    <div class="is-block">
      {outputs.map(output => {
        return (
          <dl class="box">
            <dt class="title is-5">{output.Name}</dt>
            <dd class="">{output.Description}</dd>
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
        <div class=" is-max-desktop">
          <div class="container block is-dark">
            <h2 class="title is-2">{module.Name}</h2>
            <p class="title is-6">ID: {module.ID}</p>
            <div class="tags has-addons">
              {module.Providers.map(prov => {
                return (
                  <div class="m-1">
                    <span class="tag is-info">{prov.Source}</span>
                    <span class="tag">
                      {prov.VersionConstraints}
                    </span>
                  </div>
                );
              })}
            </div>
          </div>
          <article class="panel mt-3 m-5">
            <div>
              <button class="button m-2 is-active"
                onClick$={() => store.tab = "variables"}>
                Variables
              </button>
              <button class="button m-2"
                onClick$={() => store.tab = "outputs"}>
                Outputs
              </button>
            </div>
            <div class="box">
              <div class="is-block mb-5">
                <p class="control has-icons-left">
                  <input class="input is-link" type="text" placeholder="Search" />
                  <span class="icon is-left">
                    <i class="fas fa-search" aria-hidden="true"></i>
                  </span>
                </p>
              </div>
              {store.tab === "variables" ? variablesTab(module.Variables) : <></>}
              {store.tab === "outputs" ? outputsTab(module.Outputs) : <></>}
            </div>
          </article>
        </div>
      )}
    />
  );
});
