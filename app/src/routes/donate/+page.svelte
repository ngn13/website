<script>
  import Header from "../../lib/header.svelte";
  import Card from "../../lib/card.svelte";
</script>

<svelte:head>
  <title>[ngn] | donate</title> 
  <meta content="[ngn] | donate" property="og:title" />
  <meta content="Give me all of your life savings" property="og:description" />
  <meta content="https://ngn.tf" property="og:url" />
  <meta content="#000000" data-react-helmet="true" name="theme-color" />
</svelte:head>

<Header>
  <c>bash</c> 
  donate.sh 
</Header>

<main>
  <Card title="bash donate.sh">
    If you want more services or if you want to support my work, feel free to donate. I pay around 300 Turkish Liras per month 
    for all the hosting stuff, so even a small amount will be appreciated! You can send donations to the following XMR address:
    <br>
    <br>
    <code>
      46q7G7u7cmASvJm7AmrhmNg6ctS77mYMmDAy1QxpDn5w57xV3GUY5za4ZPZHAjqaXdfS5YRWm4AVj5UArLDA1retRkJp47F 
    </code>
  </Card>
</main>

<style>
main{
  display: flex;
  flex-direction: column;
  gap: 35px;
  padding: 50px;
}

code {
  background: var(--dark-two);
  color: white;
  border-radius: var(--radius);
  text-shadow: var(--text-shadow);
  word-wrap: break-word;
  white-space: pre-wrap;
  word-break: break-word;
}
</style>
