import { component$ } from '@builder.io/qwik';

export  interface MenuItem {
  Link: string
  Label: string
}

export default component$(( props: { items: MenuItem[] }) => {
  return (
    <nav class="navbar is-transparent is-dark" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="/">
          <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28" />
        </a>

        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbar">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbar" class="navbar-menu">
        <div class="navbar-start">
        {props.items.map(item => {
          return (
            <a class="navbar-item" href={item.Link}>
              {item.Label}
            </a>
        )
        })}
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <div class="buttons">
              <a class="button is-primary">
                <strong>Sign up</strong>
              </a>
              <a class="button is-success">
                Log in
              </a>
            </div>
          </div>
        </div>
      </div>
    </nav>
    )
})