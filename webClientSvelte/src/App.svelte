<script lang="ts">
  import { Route, Router } from 'svelte-routing';
  import { authToken, pageTitle } from './api/stores';
  import AuthCallback from './components/auth/AuthCallback.svelte';
  import ServiceCallback from './components/auth/ServiceCallback.svelte';
  import InfoCallback from './components/auth/InfoCallback.svelte';
  import HomeNavbar from './components/common/HomeNavbar.svelte';
  import Home from './routes/Home.svelte';
  import Pricing from './routes/Pricing.svelte';
  import Profile from './routes/Profile.svelte';
  import Services from './routes/Services.svelte';
  import Showcase from './routes/Showcase.svelte';
  import Signin from './routes/Signin.svelte';
  import Apk from './components/common/Apk.svelte';

  export const url = "";
  let isAuth: boolean = false;
  let title: string | null = null;

  authToken.subscribe(value => {
    isAuth = value !== null;
    if (value) {
      localStorage.setItem("authToken", value);
    } else {
      localStorage.removeItem("authToken")
    }
  });

  pageTitle.subscribe(value => {
    title = value ;
  })


</script>

<Router url={url}>
  {#if isAuth}
    <main>
      <HomeNavbar>
        <h1 class="title">{title}</h1>
      </HomeNavbar>
      <Route path="*" component={Home} />
      <Route path="/home" component={Home} />
      <Route path="/services" component={Services} />
      <Route path="/profile" component={Profile} />
      <Route path="/services/auth" component={ServiceCallback} />
      <Route path="/auth/info" component={InfoCallback} />
    </main>
  {:else}
    <main>
      <Route path="*" component={Showcase} />
      <Route path='/' component={Showcase} />
      <Route path="/pricing" component={Pricing} />
      <Route path="/signin" component={Signin} />
      <Route path="/oauth" component={AuthCallback} />
      <Route path="/client.apk" component={Apk} />
    </main>
  {/if}
</Router>

<style>
  main {
    width: 100%;
    height: 100vh;
    margin: 0;
    padding: 0;
  }

  .title {
        font-size: 4rem;
        transform: translateY(5%);
  }

  @media (max-width: 768px) {
    .title {
        display: none;
    }
  }
</style>
