<script lang="ts">
    import { createEventDispatcher, onMount } from 'svelte';
    import Placeholder from 'virtual:icons/charm/plus';
    import FacebookLogo from 'virtual:icons/devicon/facebook';
    import GithubLogo from 'virtual:icons/devicon/github';
    import LinkedinLogo from 'virtual:icons/devicon/linkedin';
    import TwitterLogo from 'virtual:icons/devicon/twitter';
    import DeleteLogo from 'virtual:icons/ep/close-bold';
    import HttpLogo from 'virtual:icons/iconoir/internet';
    import DropboxLogo from 'virtual:icons/logos/dropbox';
    import EthereumLogo from 'virtual:icons/logos/ethereum';
    import GmailLogo from 'virtual:icons/logos/google-gmail';
    import NotionLogo from 'virtual:icons/logos/notion-icon';
    import SlackLogo from 'virtual:icons/logos/slack-icon';
    import SpotifyLogo from 'virtual:icons/logos/spotify-icon';
    import TwitchLogo from 'virtual:icons/logos/twitch';
    import MiroLogo from 'virtual:icons/simple-icons/miro';
    import DiscordLogo from 'virtual:icons/skill-icons/discord';
    import { Action, Reaction } from '../../api/stores';
    const dispatch = createEventDispatcher();

    let actionName: string | null = null;
    let actionRoute: string | null = null;
    export let actionType: "action" | "reaction";

    const servicesLogos = [
        { name: 'discord', icon: DiscordLogo },
        { name: 'github', icon: GithubLogo },
        { name: 'gmail', icon: GmailLogo },
        { name: 'facebook', icon: FacebookLogo },
        { name: 'spotify', icon: SpotifyLogo },
        { name: 'miro', icon: MiroLogo },
        { name: 'twitter', icon: TwitterLogo },
        { name: 'twitch', icon: TwitchLogo },
        { name: 'notion', icon: NotionLogo },
        { name: 'slack', icon: SlackLogo },
        { name: 'dropbox', icon: DropboxLogo },
        { name: 'linkedin', icon: LinkedinLogo },
        { name: 'http', icon: HttpLogo },
        { name: 'ethereum', icon: EthereumLogo },
    ];

    function getIconComponent(routeName: string) {
        const service = servicesLogos.find(service => service.name === routeName);
        return service ? service.icon : null;
    }

    function handleClick() {
        dispatch("expand");
    }

    function deleteAction() {
        dispatch("delete");
    }

    onMount(() => {
        if (actionType === "action") {
            Action.subscribe(value => {
                if (value?.service) {
                    actionName = value.service;
                    actionRoute = value?.route;
                }
            });
        } else if (actionType === "reaction") {
            Reaction.subscribe(value => {
                if (value?.service) {
                    actionName = value.service;
                    actionRoute = value?.route;
                }
            });
        }
    });

</script>

<div class="action {$$props.class}">
    {#if actionName}
        <svelte:component this={getIconComponent(actionName)} style="font-size: 1.3rem; color: black;" />
        <span class="action-name">{actionName}</span>
        <span class="action-route">{actionRoute}</span>
        <button class="delete" on:click={deleteAction}>
            <DeleteLogo style="margin-top: 5px"/>
        </button>
    {:else}
        <button class="add" on:click={handleClick}>
            <Placeholder style="margin-left: 10px; margin-top: 10px"/>
        </button>
    {/if}
</div>

<style>
    .action {
        display: flex;
        align-items: center;
        justify-content: space-around;
        min-width: 200px;
        max-width: 300px;
        width: 80%;
        height: 50px;
        min-height: 50px;
        margin-right: 10px;
        font-size: 1.5rem;
        overflow: hidden;
        background-color: var(--color2);
        border-radius: var(--borderRadius);
    }

    .action .delete {
        background: transparent;
        border: none;
        color: white;
        font-size: 1.5rem;
        border-radius: 50%;
        cursor: pointer;
        align-items: center;
        justify-content: center;
    }

    .action .delete:hover {
        background-color: white;
        color: var(--backgroundColor);
    }

    .action .add {
        border: none;
        background: transparent;
        color: white;
        font-size:2rem;
        transition: transform 0.2s ease-in-out;
    }

    .action .add:hover {
        transform: scale(1.5);
    }
</style>
