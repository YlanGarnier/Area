<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import { deleteArea, getAreas } from "../api/auth";
    import { pageTitle } from "../api/stores";
    import Area from "../components/common/Area.svelte";
    import InfoModal from "../components/common/InfoModal.svelte";
    import NewAreaModal from "../components/common/NewAreaModal.svelte";

    interface AreaInterface {
        action_service: string,
        id: number,
        name: string,
        reaction_service: string,
        route_action_service: string,
        route_reaction_service: string
    }

    let showModal = false;
    let showNewAreaModal = false;
    let areas: AreaInterface[];
    let isSmallScreen = false;
    let isMediumScreen = false;

    onMount(async () => {
        pageTitle.set("Your Areas");
        await retrieveAreas();
        window.addEventListener('resize', checkScreenSize);
    });

    // afterUpdate(async () => {
    // });

    function closeModal() {
        showModal = false;
    }

    function closeNewAreaModal() {
        showNewAreaModal = false;
    }

    async function handleShowModal() {
        showModal = true;
    }

    async function handleShowNewAreaModal() {
        showNewAreaModal = true;
    }

    async function retrieveAreas() {
        try {
            areas = await getAreas();
        } catch (error) {
            console.error("There was an error retrieving the areas", error);
        }
    }

    async function deleteThisArea(id: number) {
        try {
            await deleteArea(id);
        } catch (error) {
            console.error("There was an error deleting the area", error);
        }
    };

    function checkScreenSize() {
        isSmallScreen = window.innerWidth < 768;
        isMediumScreen = window.innerWidth < 1200;
    }

    checkScreenSize();

    onDestroy(() => {
        window.removeEventListener('resize', checkScreenSize);
    });

</script>

<div class='showcase-container'>
    <div class="area-container">
        <div class="board">
            <input class="search-area" placeholder="Search..." />
            <div class="example-line">
                <div class="new-area">
                    <button class="add-area" on:click={handleShowNewAreaModal}>+</button>
                    <h1>New Area</h1>
                </div>
                {#if isSmallScreen}
                    <h2>Area Infos</h2>
                {:else if isMediumScreen}
                    <div class="informations">
                        <h2>Area Name</h2>
                        <h2>Area Infos</h2>
                    </div>
                {:else}
                    <div class="informations">
                        <h2>Area Name</h2>
                        <h2>Delete Area</h2>
                    </div>
                {/if}
            </div>
            <div class="separator"></div>
            {#if areas && areas.length > 0}
                {#each areas as area}
                    <Area on:delete={() => deleteThisArea(area.id)} actionRoute={area.route_action_service} reactionRoute={area.route_reaction_service} areaID={area.id} action={area.action_service} reaction={area.reaction_service} areaName={area.name} on:showModal={handleShowModal}/>
                {/each}
            {/if}
        </div>
    </div>
</div>

<InfoModal show={showModal} on:close={closeModal} />
<NewAreaModal show={showNewAreaModal} on:close={closeNewAreaModal} />

<style>
    .showcase-container {
        position: relative;
        width: 100%;
        height: 85%;
        background-color: var(--backgroundColor);
        justify-content: center;
        align-items: center;
        text-align: center;
    }

    .area-container {
        width: 100%;
        height: 100%;
        display: flex;
    }

    .area-container .board {
        width: 80%;
        transform: translateX(10%);
        height: 100%;
        display: flex;
        flex-direction: column;
        margin: 20px 40px 20px 20px;
        gap: 1rem;

    }

    .area-container .search-area {
        width: 40%;
        height: 3rem;
        background-color: transparent;
        border: 1px solid white;
        border-radius: var(--borderRadius);

        font-family: 'Lekton';
        color: white;
        font-size: 2rem;
        padding-left: 1rem;
    }

    .area-container .search-area::placeholder {
        color: white;
        font-size: 2rem;
    }

    .area-container .search-area:focus {
        outline: none;
        font-family: 'koulen';
    }

    .area-container .search-area:focus::placeholder {
        opacity: 0;
    }

    .area-container .example-line {
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 50px;
        margin-top: 20px;
    }

    .area-container .example-line .new-area {
        display: flex;
        align-items: center;
        gap: 30px;
    }

    .area-container .example-line .add-area {
        width: 50px;
        height: 50px;
        font-family: 'koulen';
        font-size: 1.8rem;
        border-radius: var(--borderRadius);
        background-color: white;
        color: var(--color2);
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease-in-out, transform 0.3s ease-in-out, color 0.3s ease-in-out;
    }

    .area-container .example-line .add-area:hover {
        transform: scale(1.2);
        background-color: var(--color3);
        color: white;
    }

    .area-container .example-line .informations {
        display: flex;
        justify-content: space-between;
        width: 42%;
    }

    .area-container .separator {
        height: 2px;
        width: 100%;
        background-color: var(--color3);
    }

</style>
