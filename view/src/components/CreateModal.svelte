<script>
  import { closeModal } from 'svelte-modals';

  // provided by Modals
  export let isOpen;

  export let title;

  export let redirect = '';
  export let urlshort = '';
  export let random = false;
  let data = {
    redirect: redirect,
    urlshort: urlshort,
    random: random,
  };

  async function create() {
    await fetch('http://127.0.0.1:3000/short', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    }).then((res) => {
      console.log(res);
      if (res.status === 200) {
        closeModal();
      }
    });
  }
</script>

{#if isOpen}
  <div role="dialog" class="modal">
    <div class="contents">
      <h2>{title}</h2>
      <form on:submit|preventDefault={create}>
        <input
          type="text"
          name="urlshort"
          placeholder="URL"
          bind:value={data.urlshort}
        />
        <input
          type="text"
          name="redirect"
          placeholder="redirect"
          bind:value={data.redirect}
        />
        <div class="checkbox-input">
          <label for="random">Random URL</label>
          <input type="checkbox" name="random" bind:checked={data.random} />
        </div>
        <button type="submit">Create</button>
      </form>
    </div>
  </div>
{/if}

<style>
  .modal {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    display: flex;
    justify-content: center;
    align-items: center;

    /* allow click-through to backdrop */
    pointer-events: none;
  }

  .contents {
    background-color: white;
    padding: 5em;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    pointer-events: auto;
    border-radius: 2em;
  }

  h2 {
    text-align: center;
    font-size: 24px;
  }

  form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
  }

  input[type='text'] {
    width: 100%;
    padding: 0.5em;
    margin-bottom: 0.5em;
  }

  .checkbox-input {
    display: flex;
    flex-direction: row-reverse;
    justify-content: space-between;
    align-items: center;
    width: 100%;
  }
</style>
