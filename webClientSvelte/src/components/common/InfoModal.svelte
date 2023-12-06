<script lang="ts">
    import { createEventDispatcher, onMount } from 'svelte';
    import CloseLogo from 'virtual:icons/ep/close-bold';
    import { deleteArea } from '../../api/auth';
    import { AreaStorage } from '../../api/stores';
    const dispatch = createEventDispatcher();

    export let show = false;
    let modal;
    let title: string;
    let id: number;
    let actionService: string;
    let reactionService: string;
    let actionRoute: string;
    let reactionRoute: string;

    function close() {
        dispatch('close');
    }

    function handleKeydown(event: KeyboardEvent) {
        if (event.key === "Enter" || event.key === "Escape") {
            close();
        }
    }

    async function deleteThisArea() {
        try {
            await deleteArea(id);
            close();
        } catch (error) {
            console.error("There was an error deleting the area", error);
        }
    };

    onMount(() => {
        AreaStorage.subscribe(value => {
            if (value) {
                title = value.name;
                actionService = value.action_service;
                reactionService = value.reaction_service;
                id = value.id;
                actionRoute = value.route_action_service;
                reactionRoute = value.route_reaction_service;
            }
        })
    })
</script>

<div class="modal" bind:this={modal} class:active={show}>
    <div role="button" tabindex="0" class="overlay" on:click={close} on:keydown={handleKeydown}></div>
    <div class="content">
        <div class="header">
            <h1 class="title">{title}</h1>
            <button class="close" on:click={close}>
                <CloseLogo />
            </button>
        </div>
        <div class="Infos">
            <h3>Action service: {actionService}</h3>
            <h3>Action route: {actionRoute}</h3>
            <h3>Reaction service: {reactionService}</h3>
            <h3>Reaction route: {reactionRoute}</h3>
        </div>
        <button class="delete" on:click={deleteThisArea}>Delete the Area</button>
    </div>
</div>

<style>
    .modal {
        display: none;
        position: absolute;
        z-index: 1;
    }
    .modal.active {
        display: block;
    }
    .overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background: rgba(0, 0, 0, 0.5);
    }
    .content {
        position: fixed;
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
        width: 50%;
        height: 50%;
        top: 50%;
        left: 25%;
        transform: translateY(-50%);
        background-color: var(--color6);
        color: white;
        padding: 2rem;
        border-radius: var(--borderRadius);
        transition: transform 0.3s ease-in-out;
    }

    .header {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
    }

    .title {
        flex-grow: 1;
        text-align: center;
        margin-left: 1rem;
    }

    .close {
        background-color: transparent;
        border: none;
        font-size: 2rem;
        color: white;
        cursor: pointer;
        transform: translateY(-40%);
    }

    .content .delete {
        border: none;
        padding: 0.5rem;
        border: 3px solid white;
        border-radius: var(--borderRadius);
        color: white;
        margin-top: 1rem;
        font-size: 1.5rem;
        cursor: pointer;
        background-color: var(--color5);
    }
</style>
