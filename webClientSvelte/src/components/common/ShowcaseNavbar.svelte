<script>
    import { Link } from "svelte-routing";
    import { slide } from 'svelte/transition';
    import MenuLogo from 'virtual:icons/mi/menu';

    let isMenuOpen = false;

    function toggleMenu() {
        isMenuOpen = !isMenuOpen;
    }
</script>

<div class="navbar">
    <nav class="logo">
        <Link to="/">Area</Link>
    </nav>
    <button class="menu-btn" on:click={toggleMenu}>
        <MenuLogo />
    </button>
    {#if isMenuOpen}
        <nav class='overlay-menu' in:slide={{delay: 0, duration: 300}} out:slide={{delay: 0, duration: 300}}>
            <Link to="/pricing">Plans</Link>
            <Link to="/signin">Sign in</Link>
            <Link to="/signin#signup">Sign up</Link>
            <Link to="/client.apk"> Download App </Link>
        </nav>
    {:else}
        <nav class='buttons'>
            <Link to="/pricing">Plans</Link>
            <Link to="/signin">Sign in</Link>
            <Link to="/signin#signup">Sign up</Link>
            <Link to="/client.apk"> Download App </Link>
        </nav>
    {/if}
</div>

<style>
    .navbar {
        position: fixed;
        z-index: 1;
        width: 90vw;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 2rem 5rem;
        font-size: 2rem;
        height: 50px;
    }

    .logo {
        justify-content: flex-start;
    }

    .buttons {
        display: flex;
        gap: 10rem;
    }

    .buttons > :global(a) {
        font-family: 'Lekton';
        text-decoration: none;
        color: white;
        transition: transform 0.3s ease-in-out;
    }

    .buttons > :global(a:hover) {
        transform: scale(1.25);
    }

    .menu-btn {
        border: none;
        background: transparent;
        z-index: 1;
        font-size: 2.5rem;
        color: white;
        cursor: pointer;
        display: none;
    }

    .overlay-menu {
        display: flex;
        flex-direction: column;
        position: absolute;
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

    @media (max-width: 1200px) {
        .buttons {
            gap: 3rem;
        }
    }

    @media (max-width: 768px) {
        .menu-btn {
            display: block;
        }

        .buttons {
            display: none;
        }
    }
</style>
