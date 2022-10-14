import { component$, Slot } from '@builder.io/qwik';
import Header from '../components/header/header';

export default component$(() => {
  return (
    <>
      <main>
        <Header />
        <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
          <div class="px-4 py-6 sm:px-0">
            <Slot/>
          </div>
        </div>
      </main>
    </>
  );
});
