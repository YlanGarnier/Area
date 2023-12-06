<script lang="ts">
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';
    import { serviceConnect } from '../../api/auth';
    import { servicesStorage } from '../../api/stores';

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const code: string | null = urlParams.get('code');
        const state: string | null = urlParams.get('state');

        if (!code || !state) {
            console.error("No code or state provided.");
            navigate('/services')
            return;
        }

        const plaform: string = 'web';
        const provider: string = state;

        try {
            await serviceConnect(code, plaform, provider);
            servicesStorage.set({success: state, failed: null});
            navigate('/services')
        } catch (error) {
            servicesStorage.set({success: null, failed: state});
            navigate('/services')
        }
    });
  </script>
