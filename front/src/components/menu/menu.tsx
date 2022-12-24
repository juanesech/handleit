import { component$ } from '@builder.io/qwik';

export  interface MenuItem {
  Link: string
  Label: string
}

export default component$(( props: { items: MenuItem[] }) => {
  return (
    <div class="box">
    <aside class="menu">
      <ul class="menu-list">
        {props.items.map(item => {
          return (
          <li>
            <a href={item.Link}>{item.Label}</a>
          </li>
        )
        })}
      </ul>
    </aside>
    </div>
  );
});
