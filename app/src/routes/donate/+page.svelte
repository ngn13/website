<script>
  import Header from "$lib/header.svelte";
  import Head from "$lib/head.svelte";
  import Icon from "$lib/icon.svelte";

  import { color } from "$lib/util.js";
  import { _ } from "svelte-i18n";
</script>

<Head title="donate" desc="give me all of your life savings" />
<Header picture="money" title={$_("donate.title")} />

<main>
  <span> </span>
  <span>
    {$_("donate.info")}
    {$_("donate.price")}
  </span>
  <br />
  <br />
  <span>
    {$_("donate.details")}
  </span>
  <table>
    <thead>
      <tr>
        <th style="color: var(--{color()})">{$_("donate.table.platform")}</th>
        <th style="color: var(--{color()})">{$_("donate.table.address")}</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>
          <Icon icon="nf-fa-monero" />
          Monero (XMR)
        </td>
        <td>
          <code>
            46q7G7u7cmASvJm7AmrhmNg6ctS77mYMmDAy1QxpDn5w57xV3GUY5za4ZPZHAjqaXdfS5YRWm4AVj5UArLDA1retRkJp47F
          </code>
        </td>
      </tr>
    </tbody>
  </table>
  <span>
    {$_("donate.thanks")}
  </span>
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
    font-size: var(--size-4);
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
