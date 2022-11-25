<template>
    <div class="h-screen w-screen flex flex-col justify-center items-center bg-primary-light">
        <h1 class="text-2xl font-bold mb-3">Log In</h1>
        <form action="/auth/login" method="post">
            <input type="email" required name="email" placeholder="E-Mail Address" class="block rounded-lg p-2 m-2" />
            <input type="password" required name="password" placeholder="Password" class="block rounded-lg p-2 m-2" />
            <div class="p-2 flex flex-row justify-center">
                <button type="submit" class="bg-primary-dark rounded px-2 m-2 hover:shadow transition">Log In</button>
                <button type="submit" class="bg-white rounded px-2 m-2 hover:shadow transition" formaction="/auth/create-account">Create Account</button>
            </div>
        </form>
    </div>
</template>