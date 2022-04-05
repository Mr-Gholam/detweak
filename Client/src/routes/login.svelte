<main class=" flex justify-center m-auto py-8 my-16  items-center">
    <form action="/login" method="post" class="flex flex-col justify-between items-center  mr-56  p-4 gap-4">
       <section class="w-80">
           <label class="text-base" for="email">Email</label>
           <input type="email" name="email" id="email" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-5"  placeholder="Email" >
       </section>
       <section class="w-80">
           <label class="text-lg" for="password">Password</label>
           <input type="password" name="password" id="password" class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-5" placeholder="Password" >
       </section>
       <section class="my-3">
            <h4 class="text-base">Forgot your Password? <a href="/reset-password" class="text-lg text-main ">Reset Password</a></h4>
       </section>
       <section class="w-80">
           <input type="submit" value="Login" class="cursor-pointer text-lg rounded-lg bg-main py-2 px-20 hover:shadow-xl mx-auto block w-11/12 hover:bg-main-darker">
       </section>
       <h4 class="text-base my-5">Not a member yet?<a href="/signup" class="text-lg text-main ">Sign up</a></h4>
   </form>
   <h1 class="mr-56 p-4">there will be a pic here</h1>
</main>