<script>
  import Header from "$lib/header.svelte";
  import Head from "$lib/head.svelte";
  import { color } from "$lib/util.js";
</script>

<Head title="donate" desc="give me all of your life savings" />
<Header title="donate money!" picture="money" />

<main>
  <span
    >I spend a lot of time working on different projects and maintaining different services.</span
  >
  <span
    >I also spend a lot of money, but unlike time, you don't usually get much of it for free.</span
  >
  <span
    >I mostly pay for hosting and electricity. Which when added up costs around 550₺ per month, that
    is Turkish Lira, equals to ~$15 at time of writing.</span
  >
  <br />
  <br />
  <span
    >So even a small donation would be highly appreciated and it would help me keep everything up
    and running.</span
  >
  <table>
    <thead>
      <tr>
        <th style="color: var(--{color()})">Platform</th>
        <th style="color: var(--{color()})">Address/Link</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>Monero (XMR)</td>
        <td>
          <code>
            46q7G7u7cmASvJm7AmrhmNg6ctS77mYMmDAy1QxpDn5w57xV3GUY5za4ZPZHAjqaXdfS5YRWm4AVj5UArLDA1retRkJp47F
          </code>
        </td>
      </tr>
    </tbody>
  </table>
  <span
    >Also huge thanks to all of you who has donated so far, as I said, I highly appreciate it. Thank
    you!</span
  >
</main>

<style>
  main {
    padding: 50px;
  }

  main span {
    font-size: var(--size-4);
    color: var(--white-1);
  }

  table {
    box-shadow: var(--box-shadow);
    background: var(--black-3);
    border-collapse: collapse;
    font-size: var(--size-3);
    margin: 30px 0 30px 0;
    width: 100%;
  }

  tr,
  th,
  td {
    color: white;
    background: var(--dark-two);
    text-align: left;
  }

  td,
  th {
    border: solid 1px var(--black-4);
    padding: 16px;
  }

  th {
    font-weight: 1000;
  }

  td {
    color: var(--white-2);
    font-weight: 400;
  }

  code {
    word-wrap: break-word;
    white-space: pre-wrap;
    word-break: break-word;
  }
</style>
