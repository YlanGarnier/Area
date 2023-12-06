<script lang="ts">
    import gsap from 'gsap';
    import { ScrollTrigger } from 'gsap/ScrollTrigger';
    import { createEventDispatcher } from 'svelte';
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
    import MiroLogo from 'virtual:icons/simple-icons/miro';
    import DiscordLogo from 'virtual:icons/skill-icons/discord';
    import { newArea } from '../../api/auth';
    import { Action as ActionStore, Reaction as ReactionStore } from '../../api/stores';
    import services from "../../assets/services.json";
    import Action from './Action.svelte';
    const dispatch = createEventDispatcher();
    gsap.registerPlugin(ScrollTrigger);

    type SelectedServiceType = {
        name: string;
        route: string;
        params: any;
        ready: boolean;
    };

    let actions = services.actions;
    let reactions = services.reactions;
    let routes: any;
    let params: any;

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

    let selectedService: SelectedServiceType = {
        name: "",
        route: "",
        params: [],
        ready: false
    };

    let selectedAction: SelectedServiceType = {
        name: "",
        route: "",
        params: [],
        ready: false
    };

    let selectedReaction: SelectedServiceType = {
        name: "",
        route: "",
        params: [],
        ready: false
    };

    export let show = false;
    let modal;
    let title: string;
    let isExpanded = false;
    let serviceType: "action" | "reaction" = "action";
    let chooseRoutes: boolean = false;
    let paramsModalOpen = false;
    let reMountAction: boolean = true;
    let reMountReaction: boolean = true;
    let saveErrorMessage: string = "";

    function getIconComponent(routeName: string) {
        const service = servicesLogos.find(service => service.name === routeName);
        return service ? service.icon : null;
    }

    function cleanStores() {
        ActionStore.set(null);
        ReactionStore.set(null);
    }

    function deleteAction() {
        ActionStore.set(null);
        selectedAction.ready = false;
        reMountAction = false;
        setTimeout(() => reMountAction = true, 0);
    }

    function deleteReaction() {
        ReactionStore.set(null);
        selectedReaction.ready = false;
        reMountReaction = false;
        setTimeout(() => reMountReaction = true, 0);
    }

    function close() {
        dispatch('close');
        isExpanded = false;
        cleanStores();
        saveErrorMessage = "";
    }

    function expandModal() {
        chooseRoutes = false;
        paramsModalOpen = false;
        isExpanded = true;
    }

    function handleKeydown(event: KeyboardEvent) {
        if (event.key === "Enter" || event.key === "Escape") {
            close();
        }
    }

    function saveArea() {
        if (!selectedReaction.ready || !selectedAction.ready || !title) {
            saveErrorMessage = "Area incomplete"
            return;
        } else {
            newArea(selectedAction, selectedReaction, title);
            close();
        }
    }

    function combineParams(paramsArray: any) {
        return paramsArray.reduce((accumulator: any, currentValue: any) => {
            return {...accumulator, ...currentValue};
        }, {});
    }

    function saveParams() {
        isExpanded = false;
        if (serviceType === 'action') {
            selectedAction = { ...selectedService };
            selectedAction.ready = true;
            selectedAction.params = combineParams(selectedAction.params)
            ActionStore.set({
                service: selectedAction.name,
                route: selectedAction.route
            });
        } else if (serviceType === 'reaction') {
            selectedReaction = { ...selectedService };
            selectedReaction.ready = true;
            ReactionStore.set({
                service: selectedReaction.name,
                route: selectedReaction.route
            });
        }
        selectedService.params = [];
    }

    function selectRoute(route: any) {
        selectedService.route = route.name;
        if (serviceType === "action") {
            params = route.params;
            if (Object.keys(params).length != 0) {
                paramsModalOpen = true;
                selectedService.params = Object.keys(params).map(param => ({ [param]: "" }))
            } else
                saveParams();
        } else if (serviceType === "reaction") {
            params = route;
            if (params.target) {
                paramsModalOpen = true;
            } else
                saveParams();
        }
        chooseRoutes = false;
    }

    function selectService(service: any) {
        selectedService.name = service.name;
        routes = service.routes;
        chooseRoutes = true;
    }
</script>

<div class="modal" bind:this={modal} class:active={show}>
    <div role="button" tabindex="0" class="overlay" on:click={close} on:keydown={handleKeydown}></div>
    <div class="services-column" class:expanded={isExpanded}>
        {#if paramsModalOpen}
            <div class="params-modal">
                <h2>Enter Parameters for {selectedService.name}</h2>
                {#if serviceType === "action"}
                    {#each selectedService.params as paramObj, index (index)}
                        <div class="param-container">
                            {#each Object.keys(paramObj) as key}
                                <label for="{key}">{key}</label>
                                <input type="text" id="{key}" bind:value={selectedService.params[index][key]} />
                            {/each}
                        </div>
                    {/each}
                {:else if serviceType === "reaction"}
                    <div class="param-container">
                        <label for={params}>target</label>
                        <input type="text" bind:value={selectedService.params[0]}>
                    </div>
                {/if}
                <button class="save-action" on:click={() => { saveParams(); paramsModalOpen = false}}>
                    {#if serviceType === 'action'}
                        Save Action
                    {:else if serviceType === 'reaction'}
                        Save Reaction
                    {/if}
                </button>
            </div>
        {:else}
            {#if chooseRoutes}
                {#each routes as route}
                    <button class="route" on:click={() => selectRoute(route)}>
                        <svelte:component this={getIconComponent(selectedService.name)} style="font-size: 1.3rem; color: black;" />
                        {route.name}
                    </button>
                {/each}
            {:else if serviceType === 'action'}
                {#each actions as service}
                    <button class="route" on:click={() => selectService(service)}>
                        <svelte:component this={getIconComponent(service.name)} style="font-size: 1.3rem; color: black;" />
                        {service.name}
                    </button>
                {/each}
            {:else if serviceType === 'reaction'}
                {#each reactions as service}
                    <button class="route" on:click={() => selectService(service)}>
                        <svelte:component this={getIconComponent(service.name)} style="font-size: 1.3rem; color: black;" />
                        {service.name}
                    </button>
                {/each}
            {/if}
        {/if}
    </div>
    <div class="content" class:expanded={isExpanded}>
        <div class="header">
            <h1 class="title">Title:<input bind:value={title} type="text" placeholder="..."/></h1>
            <button class="close" on:click={close}>
                <CloseLogo />
            </button>
        </div>
        <div class="area-content">
            <div class="area-actions">
                <h1>Action</h1>
                {#if reMountAction}
                    <Action class="action" actionType="action" on:delete={deleteAction} on:expand={() => { expandModal(); serviceType = "action"; }} />
                {/if}
                <h1>Reaction</h1>
                {#if reMountReaction}
                    <Action class="reaction" actionType="reaction" on:delete={deleteReaction} on:expand={() => { expandModal(); serviceType = "reaction"; }} />
                {/if}
                <div class="options">
                    <button on:click={saveArea}>Save</button>
                    <h2>{saveErrorMessage}</h2>
                </div>
            </div>
        </div>
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

    .title input {
        border: none;
        height: 30px;
        width: 30%;
        margin-left: 10px;
        background-color: transparent;
        color: white;
        font-size: 2rem;
        font-family: 'Koulen';
    }

    .title input::placeholder {
        color: white;
        font-size: 2rem;
    }

    .title input:focus {
        outline: none;
    }

    .close {
        background-color: transparent;
        border: none;
        font-size: 2rem;
        color: white;
        cursor: pointer;
        transform: translateY(-40%);
    }

    .area-content {
        position: relative;
        display: flex;
        justify-content: space-around;
        width: 100%;
        height: 100%;
        overflow-y: hidden;
    }

    .area-actions {
        position: relative;
        text-align: center;
        justify-content: flex-start;
        justify-self: center;
        height: 100%;
        width: 100%;
        overflow: hidden;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .options {
        margin-top: 1rem;
        display: flex;
        flex-direction: column;
        justify-content: flex-end;
    }

    .options button {
        padding: 0.3rem;
        background-color: var(--color2);
        border: none;
        border: 2px solid white;
        color: white;
        font-size: 2rem;
        cursor: pointer;
        border-radius: var(--borderRadius);
    }

    .services-column {
        position: fixed;
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 15%;
        height: 50%;
        top: 50%;
        left: 60%;
        transform: translateY(-50%);
        background-color: var(--color6);
        color: white;
        padding: 2rem;
        border-radius: var(--borderRadius);
        transition: transform 0.3s ease-in-out;
        gap: 1rem;
    }

    .services-column.expanded {
        transform: translateY(-50%) translateX(40%);
    }

    .content.expanded {
        transform: translateY(-50%) translateX(-20%);
    }

    .services-column .route {
        display: flex;
        justify-content: flex-start;
        gap: 2rem;
        padding-left: 1rem;
        align-items: center;
        width: 100%;
        height: 2rem;
        border: none;
        border-radius: var(--borderRadius);
        color: white;
        font-size: 1.5rem;
        background-color: var(--color2);
        cursor: pointer;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .params-modal {
        position: fixed;
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 1rem;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        z-index: 1;
    }

    .params-modal .param-container {
        position: relative;
        color: white;
        border-radius: var(--borderRadius);
        display: flex;
        flex-direction: column;
    }

    .params-modal .param-container > label {
        font-size: 1.5rem;
    }

    .params-modal .param-container > input {
        height: 2rem;
        border-radius: var(--borderRadius);
        border: none;
        color: var(--color6);
        font-size: 1rem;
        font-weight: bold;
    }

    .params-modal .param-container > input::placeholder {
        color: var(--color6);
    }

    .params-modal .param-container > input:focus {
        outline: none;
    }

    .params-modal .save-action {
        padding: 0.5rem;
        border: none;
        background-color: var(--color2);
        color: white;
        cursor: pointer;
        font-size: 1.5rem;
        border-radius: var(--borderRadius);
    }

    @media (max-width: 1200px) {
        .content {
            left: 0;
            width: 90%;
        }

        .services-column {
            left: 0;
            width: 90%;
        }

        .content.expanded {
            transform: translateY(-50%) translateX(0);
        }

        .services-column.expanded {
            transform: translateY(-50%) translateX(0);
            z-index: 1;
        }
    }

</style>
