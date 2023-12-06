<script lang="ts">
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';

  onMount(() => {
    const pathAsset: string = 'src/assets/app-release.apk';

    const apkLink: HTMLAnchorElement = document.createElement('a');
    apkLink.href = pathAsset;
    apkLink.download = 'app-release.apk';

    document.body.appendChild(apkLink);
    apkLink.click();

    document.body.removeChild(apkLink);

    navigate('/');
  });
</script>
