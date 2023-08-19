<script>
  import { page } from "$app/stores"
  export let link 
</script>

<div>
  <a href="{link}"><slot></slot></a>
</div>

<style>
a {
  margin-left: 20px;
  font-weight: 700;
  font-size: 22px;
  text-decoration: none;
  color: white;
  cursor: pointer;
}

a:hover {
  text-decoration: underline;
  animation-name: underlineAnimation;
  animation-duration: 5s;
  animation-iteration-count: infinite;
  text-shadow: 3px 4px 7px rgba(81, 67, 21, 0.8);
}
</style>
