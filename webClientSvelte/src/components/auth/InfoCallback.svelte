<script lang="ts">
    import { navigate } from 'svelte-routing';
    import { updateInfos } from '../../api/auth';
    import { authToken } from '../../api/stores';

    let username: string = '';
    let firstName: string = '';
    let lastName: string = '';
    let code: string | null = null
    let state: string | null = null

    async function submitInfos() {

        if(!username || !firstName || !lastName)
            return;

        try {
            await updateInfos(firstName, lastName, username);
            navigate('/home');
        } catch (error) {
            authToken.set(null);
            navigate('/signin#signup');
        }
    }
</script>

<form on:submit|preventDefault={submitInfos}>
    <h1>We need some more informations about you !</h1>
    <input bind:value={username} type="text" placeholder="Username" />
    <input bind:value={firstName} type="text" placeholder="First name" />
    <input bind:value={lastName} type="text" placeholder="Last name" />
    <button class="submitbutton">Sign Up</button>
</form>

<style>
    form {
        width: 50%;
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 2rem;
        transform: translateX(50%);
        text-align: center;
        align-items: center;
    }

    input {
        padding: 1rem;
        width: 30%;
        border: none;
        border-radius: var(--borderRadius);
    }

    input:hover {
        outline: none;
    }

    button {
        padding: 0.5rem;
        border: none;
        border-radius: var(--borderRadius);
        cursor: pointer;
        color: var(--color6);
        font-size: 1.5rem;
    }
</style>
