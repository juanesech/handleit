import { component$, Slot } from '@builder.io/qwik';
import Nav from '../components/nav/nav';

export default component$(() => {
  const items = [
    {
      Label: "Modules",
      Link: "/"
    },
    {
      Label: "Settings",
      Link: "/settings"
    }
  ]
  return (
    <>
      <main>
        <Nav items={items} />
        <div class="container mt-5">
          <Slot />
        </div>
      </main>
    </>
  );
});
