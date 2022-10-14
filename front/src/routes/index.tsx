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

export const onGet: RequestHandler<ModuleSummary[]> =async () => {
  let modules:Array<ModuleSummary> = []
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
        <div class="block content-center mx-auto max-w-xs">
          {moduleList.map(module => {
              return (
                <a  href={`/modules/${module.Name}`} class="m-2 group block max-w-xs mx-auto rounded-lg p-6 bg-white ring-1 ring-slate-900/5 shadow-lg space-y-3 hover:bg-sky-500 hover:ring-sky-500">
                  <div>
                    <div class="group-hover:text-white text-xl font-medium text-black">{module.Name}</div>
                  </div>
                </a>
              )
            })}
        </div>
      )}
    />
  );
});