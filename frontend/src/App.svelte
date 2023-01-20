<script>
    import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
    import Navbar from "./components/Navbar.svelte";
    import { Route } from "tinro"
    import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import { authState } from "./api/auth";
  import SiteInfo from "./pages/SiteInfo.svelte";

    const queryClient = new QueryClient()
</script>

<div class="dark:bg-slate-900 dark:text-white">
<QueryClientProvider client={queryClient}>
    {#if $authState}
    <div>
    <Navbar/>
    <Route path="/"><Home/></Route>
    <Route path="/sites/:id" let:meta><SiteInfo siteId={meta.params.id}/></Route>
    </div>
    {:else}
    <Login/>
    {/if}
</QueryClientProvider>
</div>