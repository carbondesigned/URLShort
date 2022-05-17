<script>
  import { closeModal } from 'svelte-modals';
  export let isOpen;
  export let title;
  export let id;

  async function deleteURL() {
    await fetch(`http://127.0.0.1:3000/short/${id}`, {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
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
      <form on:submit|preventDefault={deleteURL}>
        <button onclick={() => closeModal()}>Cancel</button>
        <button type="submit">Delete</button>
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
</style>
