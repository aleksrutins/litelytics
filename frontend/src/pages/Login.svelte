<div class="h-screen w-screen flex flex-col justify-center items-center">
    <h1 class="heading">Log In</h1>
    <form action="/auth/login" method="post" class="flex flex-col items-stretch">
        <input type="email" required name="email" placeholder="E-Mail Address" class="block rounded-lg p-2 m-2 bg-slate-50 dark:bg-slate-800" />
        <input type="password" required name="password" placeholder="Password" class="block rounded-lg p-2 m-2 bg-slate-50 dark:bg-slate-800" />
        <div class="p-2 flex flex-row justify-center">
            <button type="submit" class="bg-black text-white dark:bg-slate-800 rounded px-2 m-2 hover:shadow transition">Log In</button>
            <button type="submit" class="bg-white dark:bg-black rounded px-2 m-2 hover:shadow transition" formaction="/auth/create-account">Create Account</button>
        </div>
    </form>
</div>