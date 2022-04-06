<main class="flex justify-center m-auto py-8  items-center">
    <form action="/signup" method="post" class="flex flex-col justify-between mr-56 p-4 gap-3">
        <section class="w-80">
            <label class="text-base" for="email">Email</label>
            <input type="email" name="email" id="email" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3" >
        </section>
        <section class="w-80">
            
            <label class="text-base" for="username">Username</label>
            <input type="text" name="username" id="username" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3" >
        </section>
        <section class="w-80">
            <label class="text-base" for="password">Password</label>
            <input type="password" name="password" id="password" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3" >
        </section>
        <section class="w-80">
            <label class="text-base" for="confirmPassword">Confirm Password</label>
            <input type="password" name="confirmPassword" id="confirmPassword" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3" >
        </section>
        <section class="m-2">
            <input type="checkbox" name="privacy" id="">
            <span>Privacy bullshit</span>
        </section>
        <section  class="w-80">
            <input type="submit" value="Sign Up" class="cursor-pointer text-lg rounded-lg text-white bg-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl hover:bg-main-darker  ">
        </section>
        <h5 class="text-base text-center">Already a member?<a href="/login" class="text-lg text-main mx-1">Login</a></h5>
    </form>
    <h1 class="ml-56 p-4">There will be pic here</h1>
</main>