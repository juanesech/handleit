import { component$ } from '@builder.io/qwik';

export default component$(() => {

  return (
    <div class="relative bg-white">
      <div class="mx-auto max-w-7xl px-4 sm:px-6">
        <div class="flex items-center justify-between border-b-2 border-gray-100 py-6 md:justify-start md:space-x-10">
          <div class="flex justify-start lg:w-0 lg:flex-1">
            <a href="#">
              <span class="sr-only">Topo</span>
            </a>
          </div>
          <nav class="hidden space-x-10 md:flex">
            <a href="/" class="text-base font-medium text-gray-500 hover:text-blue-700">Modules</a>
            <a href="/settings" class="text-base font-medium text-gray-500 hover:text-blue-700">Settings</a>
          </nav>
          <div class="hidden items-center justify-end md:flex md:flex-1 lg:w-0">
            <a href="#" class="ml-8 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-transparent bg-blue-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-blue-700">Sign in</a>
          </div>
        </div>
      </div>
      <div class="absolute inset-x-0 top-0 origin-top-right transform p-2 transition md:hidden">
        <div class="divide-y-2 divide-gray-50 rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5">
          <div class="px-5 pt-5 pb-6">
            <div class="flex items-center justify-between">
              <div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
});
