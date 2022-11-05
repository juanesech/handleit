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
    const store = useStore({ editMode: false })

    return (
        <Resource
            value={source}
            onPending={() => <div>Loading...</div>}
            onRejected={() => <div>Error</div>}
            onResolved={(source) => (
                <div class="block content-center mx-auto max-w-xl">
                        <div class="bg-blue-600 ring-blue-600 list-none flex flex-wrap items-center rounded-sm">
                        <div class="rounded-sm px-4 py-3 sm:px-6 group-hover:bg-blue-600 hover:ring-blue-600">
                            <h4 class="text-white text-md">{source.name}</h4>
                            <span class="align-sub text-sm font-medium inline-block py-1 px-2 rounded group-hover:text-semibold group-hover:text-blue-500 bg-blue-100 lowercase last:mr-0 mr-1">
                                {source.type}
                            </span>
                        </div>
                        <div class="hidden items-center justify-end md:flex md:flex-1 lg:w-0">
                            <a href="#" class="mx-3 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-blue-500 bg-white px-4 py-2 text-base font-light text-blue shadow-sm hover:text-white hover:bg-blue-500">Modify</a>
                        </div>
                        </div>
                        <dl>
                        <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                                <dt class="text-md font-medium text-gray-500">Id</dt>
                                <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{source.ID}</dd>
                            </div>
                            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                                <dt class="text-md font-medium text-gray-500">Type</dt>
                                <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{source.type}</dd>
                            </div>
                            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                                <dt class="text-md font-medium text-gray-500">Address</dt>
                                <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{source.address}</dd>
                            </div>
                            {source.type === "GitLab" ?
                                <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                                    <dt class="text-md font-medium text-gray-500">Group</dt>
                                    <dd class="mt-1 text-md text-gray-900 sm:col-span-2 sm:mt-0">{source.group}</dd>
                                </div> : <></>}
                        </dl>
                </div>
            )}
        />
    );
});
