<script>
  import Header from "../../lib/header.svelte";
  import Card from "../../lib/card.svelte";
</script>

<svelte:head>
  <title>[ngn.tf] | donate</title> 
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
    I work on free/libre and open source software and offer free services. General hosting 
    and stuff costs around 550₺ (~$17) per month, so feel free to donate in order to help me keep
    everything up and running!
    <table>
      <tr>
        <th>Platform</th>
        <th>Address/Link</th>
      </tr>
      <tr>
        <td>Monero (XMR)</td>
        <td>    
          <code>
          46q7G7u7cmASvJm7AmrhmNg6ctS77mYMmDAy1QxpDn5w57xV3GUY5za4ZPZHAjqaXdfS5YRWm4AVj5UArLDA1retRkJp47F 
          </code>
        </td>
      </tr>
    </table>
    Also huge thanks to all of you who has donated so far, even if it's a small amount, I highly
    appreciate it. Thank you!
  </Card>
</main>

<style>
main{
  display: flex;
  flex-direction: column;
  gap: 35px;
  padding: 50px;
}

table {
  border-collapse: collapse;
  border: none;
  color: white;
  font-size: 20px;
  width: 100%;
  margin: 30px 0 30px 0;
  box-shadow: var(--box-shadow);
}


tr,th,td{
  color: white;
  background: var(--dark-two);
}

td,th{
  border: solid 1px var(--dark-fife);
  padding: 16px;
}

th {
  animation-name: colorAnimation;
  animation-duration: 10s;
  animation-iteration-count: infinite;
  background: var(--dark-two);
}

code {
  word-wrap: break-word;
  white-space: pre-wrap;
  word-break: break-word;
}
</style>
