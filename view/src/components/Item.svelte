<script>
  import { openModal } from 'svelte-modals';
  import Card from './Card.svelte';
  import DeleteModal from './DeleteModal.svelte';
  export let short;

  function handleOpenDeleteConfirmationModal(title, id) {
    openModal(DeleteModal, { title: title, id: id });
  }
</script>

<Card>
  <p>
    Short: <a href="http://127.0.0.1:3000/r/{short.urlshort}"
      >http://127.0.0.1:3000/r/{short.urlshort}</a
    >
  </p>
  <p>Redirect: {short.redirect}</p>
  <p>Clicked: {short.clicked}</p>
  <!-- <button class="update">Update</button> -->
  <button
    class="delete"
    on:click={() =>
      handleOpenDeleteConfirmationModal(
        'Confirm Deletion of URLShort',
        short.id
      )}>Delete</button
  >
</Card>

<style>
  /*   .update {
    background-color: rgb(15, 126, 210);
    border: none;
    color: white;
    padding: 0.5em 1em;
    border-radius: 0.25em;
  } */
  .delete {
    background-color: rgb(206, 0, 0);
    border: none;
    color: white;
    padding: 0.5em 1em;
    border-radius: 0.25em;
  }
</style>
