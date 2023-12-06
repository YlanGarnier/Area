<script lang="ts">
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';
    import { get } from 'svelte/store';
    import { oAuthSignup, updateInfos } from '../../api/auth';
    import { AuthStatus, authToken } from '../../api/stores';

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const code = urlParams.get('code');
        const state = urlParams.get('state');

        if (!code || !state) {
            console.error("No code or state provided.");
            navigate('signin#signup');
            return;
        }

        const plaform: string = 'web';
        const provider: string = state;

        try {
            const token = await oAuthSignup(code, plaform, provider);
            authToken.set(token);
            const userExist = get(AuthStatus);
            console.log(userExist);
            if (userExist)
                navigate('/home');
            else
                navigate('/auth/info')
        } catch (error) {
            navigate('/signin#signup');
        }
    });
</script>

