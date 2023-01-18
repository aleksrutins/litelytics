<script lang="ts">
    import { createLink, Link } from "@bjornlu/svelte-router"
    import { getEmail } from "../api/auth";

    const home = createLink("/");
</script>
<div class="border-b flex flex-row content-between">
    <div>
        <img src="../assets/logo.webp" alt="Litelytics logo" width="100" class="self-center dark:invert"/>
        <Link to="/" class={($home.isExactActive ? "bg-slate-300" : "") + "rounded-sm hover:bg-slate-400"}>Home</Link>
    </div>
    <div>
        {getEmail()}
    </div>
</div>