<script lang="ts">
    import { createEventDispatcher, onDestroy, onMount } from 'svelte';
    import FacebookLogo from 'virtual:icons/devicon/facebook';
    import GithubLogo from 'virtual:icons/devicon/github';
    import LinkedinLogo from 'virtual:icons/devicon/linkedin';
    import TwitterLogo from 'virtual:icons/devicon/twitter';
    import CloseLogo from 'virtual:icons/ep/close-bold';
    import HttpLogo from 'virtual:icons/iconoir/internet';
    import DropboxLogo from 'virtual:icons/logos/dropbox';
    import EthereumLogo from 'virtual:icons/logos/ethereum';
    import GmailLogo from 'virtual:icons/logos/google-gmail';
    import NotionLogo from 'virtual:icons/logos/notion-icon';
    import SlackLogo from 'virtual:icons/logos/slack-icon';
    import SpotifyLogo from 'virtual:icons/logos/spotify-icon';
    import TwitchLogo from 'virtual:icons/logos/twitch';
    import ArrowLogo from 'virtual:icons/ph/arrow-right-bold';
    import MiroLogo from 'virtual:icons/simple-icons/miro';
    import DiscordLogo from 'virtual:icons/skill-icons/discord';
    import { AreaStorage } from '../../api/stores';

    const dispatch = createEventDispatcher();
    export let areaID: number;
    export let action: string;
    export let reaction: string;
    export let areaName: string;
    export let actionRoute: string;
    export let reactionRoute: string;
    let isSmallScreen = false;

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

    function showModal() {
        AreaStorage.set({
            action_service: action,
            id: areaID,
            name: areaName,
            reaction_service: reaction,
            route_action_service: actionRoute,
            route_reaction_service: reactionRoute
        });
        dispatch('showModal');
    }

    function deleteArea() {
        dispatch('delete', areaID);
    }

    function handleKeydown(event: KeyboardEvent) {
        if (event.key === "Enter") {
            showModal();
        }
    }

    function getIconComponent(routeName: string) {
        const service = servicesLogos.find(service => service.name === routeName);
        return service ? service.icon : null;
    }

    function checkScreenSize() {
        isSmallScreen = window.innerWidth < 1200;
    }

    checkScreenSize();

    onMount(() => {
        window.addEventListener('resize', checkScreenSize);
    });

    onDestroy(() => {
        window.removeEventListener('resize', checkScreenSize);
    });
</script>


<div class="area">
    <div class="services">
        <svelte:component this={getIconComponent(action)} style="font-size: 2rem;" />
        <span class="service-text">{action}</span>
        {#if !isSmallScreen} <ArrowLogo style="font-size: 2rem;" /> {/if}
        <svelte:component this={getIconComponent(reaction)} style="font-size: 2rem;" />
        <span class="service-text">{reaction}</span>
    </div>
    <h1 class="area-name">{areaName}</h1>
    <div role="button" tabindex="0" on:click={isSmallScreen ? showModal : deleteArea} on:keydown={handleKeydown}>
        {#if isSmallScreen}
            <ArrowLogo style="font-size: 2rem; transform: translateY(3px);" />
        {:else}
            <CloseLogo style="font-size: 2rem; transform: translateY(3px);" />
        {/if}
    </div>
</div>

<style>
    .area {
        position: relative;
        background-color: white;
        isolation: isolate;
        color: var(--color2);
        height: 5rem;
        width: 100%;
        max-width: 100%;
        border-radius: var(--borderRadius);
        display: flex;
        align-items: center;
        justify-content: space-between;
        overflow: hidden;
        transition: color 0.3s ease-in-out;
        padding: 0 1rem; /* Added padding for smaller screens */
    }

    .area .services {
        margin-left: 2rem;
        display: flex;
        justify-content: space-between;
        align-items: center; /* Ensure icons are centered */
        flex-grow: 1; /* Allow this to grow and take available space */
    }

    .area .services span {
        font-size: 1.7rem;
        transform: translateY(2px);
        transition: opacity 0.3s ease-in-out;
    }

    .area-name {
        z-index: 1;
        flex-grow: 2; /* Allow the name to take more space */
        text-align: center; /* Center the name */
    }

    @media (max-width: 1200px) {
        .area .services span {
            display: none;
        }
    }

    @media (max-width: 768px) {
        .area .services {
            display: none;
        }
    }
</style>
