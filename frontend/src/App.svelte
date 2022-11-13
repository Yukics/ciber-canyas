<script>
  import { user, token, emojis } from './lib/store'
  import Footer from './components/Footer.svelte';
  import Nav from './components/Nav.svelte';
  import Elements from './components/Elements.svelte';

  //? get previous stored session values
  user.subscribe(val => sessionStorage.setItem("user", val));
  token.subscribe(val => sessionStorage.setItem("token", val));

</script>

<style>

</style>

<Nav/>  
<Elements/>
<Footer></Footer>