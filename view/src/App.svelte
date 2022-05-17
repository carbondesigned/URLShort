<script>
  import List from './components/List.svelte';
  import { Modals, closeModal, openModal } from 'svelte-modals';
  import CreateModal from './components/CreateModal.svelte';

  function handleOpenCreateURLModal(title) {
    openModal(CreateModal, { title: title });
  }
</script>

<Modals>
  <div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>
<main>
  <header>
    <h1>Sorten Your URL</h1>
    <button
      class="create"
      on:click={() => handleOpenCreateURLModal('Create a Shorter URL')}
      >Create</button
    >
  </header>
  <List />
</main>

<style>
  .backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.5);
  }
</style>
