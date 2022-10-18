import { Resource, component$, useStore } from "@builder.io/qwik";
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

export const onGet: RequestHandler<ModuleSource> = async ({ params }) => {
    try {
        const response = await axios.get(`http://localhost:8080/config/${params.source}`);
        let data = await response.data;
        console.log(data)
        return {
            ID: data.ID,
            name: data.name,
            type: data.type,
            address: data.address,
            group: data.group,
            auth: data.auth
        };
    } catch (error) {
        console.log(error);
    }
};
export default component$(() => {
  const source = useEndpoint<ModuleSource>();

  return (
    <Resource
      value={source}
      onPending={() => <div>Loading...</div>}
      onRejected={() => <div>Error</div>}
        onResolved={(source) => (
        <>
        <div class="block content-center mx-auto max-w-xl">
            <div class="hidden items-center justify-start md:flex md:flex-1 lg:w-0">
                <a href="#" class="ml-0 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-blue-500 bg-white px-4 py-2 text-base font-medium text-blue shadow-sm hover:text-white hover:bg-blue-500">Modify</a>
            </div>
            <div class="m-2 mx-auto rounded-sm p-4 bg-white ring-1">
                <ul>
                    <li>{source.ID}</li>
                    <li>{source.name}</li>
                    <li>{source.type}</li>
                    <li>{source.address}</li>
                    <li>{source.group}</li>
                    <li>{source.auth}</li>
                </ul>
            </div>
        </div>
        </>
      )}
    />
  );
});
