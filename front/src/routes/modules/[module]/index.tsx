import { Resource, component$ } from "@builder.io/qwik";
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

export const onGet: RequestHandler<Module> =async ({ params }) => {
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

  return (
    <Resource
      value={moduleData}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
      onResolved={(module) => (
        <div class="overflow-hidden bg-white shadow sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h2 class="text-xl font-medium leading-6 text-gray-900">{module.Name}</h2>
            <p class="mt-1 max-w-2xl text-sm text-gray-500">{module.ID}</p>
          </div>
            <details open>
              <summary class="list-none">
                <h2 class="cursor-pointer text-lg font-medium leading-3 sm:px-4 m-3">Variables</h2>
              </summary>
                <div aria-labelledby="headingOne"
                  data-bs-parent="#accordionExample">
                    {module.Variables.map(variable => {
                      return(
                        <div class="border-t border-gray-200">
                          <dl>
                            <h2 class="text-sm">{variable.Name}</h2>
                            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                              <dt class="text-sm font-medium text-gray-500">Description</dt>
                              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{variable.Description}</dd>
                            </div>
                            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                              <dt class="text-sm font-medium text-gray-500">Type</dt>
                              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{variable.Type}</dd>
                            </div>
                            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                              <dt class="text-sm font-medium text-gray-500">Required</dt>
                              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{String(variable.Required)}</dd>
                            </div>
                            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                              <dt class="text-sm font-medium text-gray-500">Default value</dt>
                              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{variable.Default}</dd>
                            </div>
                          </dl>
                        </div>
                      )})}
                </div>
            </details>
        </div>
      )}
    />
  );
});
