<section>
  <form method="post" action="/session/login">
    {{ .xsrfdata }}
    <div class="columns is-centered">
      <div class="column is-one-third">
        <header class="has-text-centered">
          <h1 class="title is-2">Please Login</h1>
        </header>
      </div>
    </div>
    <div class="columns is-centered">
      <div class="column is-one-third">
        <div class="card is-fullwidth">
          <div class="card-content">
            <div class="container">
              <div class="control is-horizontal">
                <div class="control-label">
                  <label class="label">Email</label>
                </div>
                <p class="control">
                  <input class="input" name="email" type="email" placeholder="Email">
                </p>
              </div>
              <div class="control is-horizontal">
                <div class="control-label">
                  <label class="label">Password</label>
                </div>
                <p class="control">
                  <input class="input" name="password" type="password" placeholder="Password">
                </p>
              </div>
              <div class="control is-clearfix">
                <button type="submit" class="button is-danger is-pulled-right">Login</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </form>
</section>