import { component$, Slot } from '@builder.io/qwik';
import Header from '../components/menu/menu';

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
        <div class="container">
        <div class="columns">
          <div class="column is-one-quarter mt-5">
            <Header items={items}/>
          </div>
          <div class="column mt-5">
            <Slot/>
          </div>
        </div>
        </div>
      </main>
    </>
  );
});
