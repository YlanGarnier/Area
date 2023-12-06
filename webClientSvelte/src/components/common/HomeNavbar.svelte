<script lang="ts">
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";
    import { slide } from 'svelte/transition';
    import MenuLogo from 'virtual:icons/mi/menu';
    import { getUserInfo, logout } from "../../api/auth";
    import { userInfos } from "../../api/stores";

    let userInitial: string | null = null;
    let userName: string | null = null;

    let isMenuOpen = false;
    let showUserModal = false;

    function openModal() {
        showUserModal = true;
    }

    function closeModal() {
        showUserModal = false;
    }

    function toggleMenu() {
        isMenuOpen = !isMenuOpen;
    }

    async function fetchUserName() {
        try {
            await getUserInfo();
            userInfos.subscribe(value => {
                if (value) {
                    userName = value.Username.String ;
                    userInitial = userName.charAt(0);
                } else {
                    userName = null;
                    userInitial = null;
                }
            });
        } catch (error: any) {
            console.error("Failed to get username:", error.message);
        }
    }

    function signOut() {
        logout();
    }

    onMount(() => {
        fetchUserName();
    });

</script>

<div class="navbar">
    <nav class="logo">
        <Link to="/home">Logo</Link>
    </nav>
    <slot></slot>
    <button class="menu-btn" on:click={toggleMenu}>
        <MenuLogo />
    </button>
    {#if isMenuOpen}
        <nav class='overlay-menu' in:slide={{delay: 0, duration: 300}} out:slide={{delay: 0, duration: 300}}>
            <Link to="/services">Services</Link>
            <Link to="/parameter">Parameter</Link>
            <Link to="/profile">Profil</Link>
        </nav>
    {:else}
        <nav class="profilButton">
            <div class="profilContainer">
                <div class="profil" class:showModal={showUserModal} on:mouseenter={openModal} on:mouseleave={closeModal} tabindex="0" role="button">
                    {#if !showUserModal}
                        {userInitial}
                    {:else}
                    <div class="usermodal-content">
                        <h1>{userName}</h1>
                        <div class="links">
                            <Link to="/profile">> Profil</Link>
                            <Link to="/home">> Areas</Link>
                            <Link to="/services">> Services</Link>
                        </div>
                        <button on:click={signOut}>Sign out</button>
                    </div>
                    {/if}
                </div>
            </div>
        </nav>
    {/if}
</div>

<style>
    .navbar {
        display: flex;
        text-align: center;
        justify-content: space-between;
        align-items: center;
        padding: 2rem 5rem;
        font-size: 2rem;
        height: 50px;
    }

    .navbar .profilButton > :global(a) {
        text-decoration: none;
        color: white;
    }

    .navbar .profilContainer {
        position: relative;
        height: 60px; /* Match the initial height of the .profil */
        width: 60px;  /* Match the initial width of the .profil */
        overflow: visible;
    }

    .navbar .profil {
        position: absolute;
        display: flex;
        z-index: 3;
        flex-direction: column;
        transform: translateX(-50%);
        border: 3px solid white;
        border-radius: 50%;
        height: 60px;
        width: 60px;
        overflow: hidden;
        transition: background-color 0.3s ease-in-out, border-color 0.3s ease-in-out, color 0.3s ease-in-out, width 0.3s ease-in-out, height 0.3s ease-in-out, border-radius 0.3s ease-in-out, transform 0.3s ease-in-out;
        justify-content: space-around;
        align-items: center;
        transform-origin: top center;
    }

    .navbar .profil.showModal {
        width: 250px;
        height: 350px;
        border-radius: 10px;
        background-color: var(--color5);
        color: white;
        transform: translateX(-90%);
    }

    .navbar .profil .usermodal-content {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .navbar .profil .usermodal-content h1 {
        text-decoration: none;
        margin-bottom: 50px;
        margin-top: 0;
        font-size: 2.5rem;
        color: white;
        transform: translateY(-10%);
    }

    .navbar .profil .usermodal-content .links {
        display: flex;
        flex-direction: column;
        align-items: start;
    }

    .navbar .profil .usermodal-content .links > :global(a) {
        text-decoration: none;
        color: white;
        font-size: 1.7rem;
    }

    .navbar .profil .usermodal-content .links > :global(a:hover) {
        text-decoration-line: underline;
    }

    .navbar .profil .usermodal-content button {
        border: none;
        margin-top: 1rem;
        padding: 0.3rem 1rem;
        background-color: transparent;
        border-radius: var(--borderRadius);
        border: 3px solid white;
        color: white;
        font-size: 1.7rem;
        font-family: 'Koulen';
        cursor: pointer;
    }

    .navbar .profil .usermodal-content button:hover {
        text-decoration-line: underline;
    }

    .navbar .menu-btn {
        border: none;
        background: transparent;
        z-index: 3;
        font-size: 2.5rem;
        color: white;
        cursor: pointer;
        display: none;
    }

    .navbar .overlay-menu {
        display: flex;
        flex-direction: column;
        position: absolute;
        z-index: 2;
        top: 0;
        right: 0;
        width: 100%;
        height: 100vh;
        background: var(--backgroundColor);
        opacity: 0.9;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    .overlay-menu > :global(a) {
        font-family: 'Lekton';
        text-decoration: none;
        color: white;
    }

    @media (max-width: 768px) {
        .menu-btn {
            display: block;
        }

        .navbar .profil {
            transform: translateX(50%);
        }

        .navbar .profil.showModal {
            transform: translateX(-50%);
        }
    }
</style>
