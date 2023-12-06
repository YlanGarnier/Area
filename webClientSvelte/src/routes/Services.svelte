<script lang="ts">
    import gsap from 'gsap';
    import { ScrollTrigger } from 'gsap/ScrollTrigger';
    import { afterUpdate, onMount } from 'svelte';
    import FacebookLogo from 'virtual:icons/devicon/facebook';
    import GithubLogo from 'virtual:icons/devicon/github';
    import LinkedinLogo from 'virtual:icons/devicon/linkedin';
    import TwitterLogo from 'virtual:icons/devicon/twitter';
    import DropboxLogo from 'virtual:icons/logos/dropbox';
    import GmailLogo from 'virtual:icons/logos/google-gmail';
    import NotionLogo from 'virtual:icons/logos/notion-icon';
    import SpotifyLogo from 'virtual:icons/logos/spotify-icon';
    import TwitchLogo from 'virtual:icons/logos/twitch';
    import MiroLogo from 'virtual:icons/simple-icons/miro';
    import DiscordLogo from 'virtual:icons/skill-icons/discord';
    import { getServices } from '../api/auth';
    import { pageTitle, servicesStorage } from '../api/stores';
    gsap.registerPlugin(ScrollTrigger);

    interface Service {
        name: string;
        icon: typeof FacebookLogo;
        status: boolean;
        url: string | null;
    }

    let services: Service[] = [
        { name: 'discord', icon: DiscordLogo, status: false,  url: `${import.meta.env.VITE_DISCORD_CLIENT_API}?client_id=${import.meta.env.VITE_DISCORD_CLIENT_ID}&scope=identify%20guilds%20email%20guilds.join%20gdm.join&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=discord` },
        { name: 'github', icon: GithubLogo, status: false,  url: `${import.meta.env.VITE_GITHUB_CLIENT_API}?client_id=${import.meta.env.VITE_GITHUB_CLIENT_ID}&scope=repo%20user&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=github`},
        { name: 'gmail', icon: GmailLogo, status: false,  url: `${import.meta.env.VITE_GOOGLE_CLIENT_API}?client_id=${import.meta.env.VITE_GOOGLE_CLIENT_ID}&access_type=offline&response_type=code&state=google_gmail&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&scope=https://www.googleapis.com/auth/userinfo.email%20https://mail.google.com/`},
        { name: 'facebook', icon: FacebookLogo, status: false,  url: `${import.meta.env.VITE_FACEBOOK_CLIENT_API}?client_id=${import.meta.env.VITE_FACEBOOK_CLIENT_ID}&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=facebook` },
        { name: 'spotify', icon: SpotifyLogo, status: false,  url: `${import.meta.env.VITE_SPOTIFY_CLIENT_API}?client_id=${import.meta.env.VITE_SPOTIFY_CLIENT_ID}&scope=user-read-email,user-read-currently-playing&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=spotify` },
        { name: 'miro', icon: MiroLogo, status: false,  url: `${import.meta.env.VITE_MIRO_CLIENT_API}?client_id=${import.meta.env.VITE_MIRO_CLIENT_ID}&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=miro`},
        { name: 'twitter', icon: TwitterLogo, status: false,  url: `${import.meta.env.VITE_TWITTER_CLIENT_API}?client_id=${import.meta.env.VITE_TWITTER_CLIENT_ID}&scope=tweet.read%20users.read%20follows.read%20follows.write&code_challenge=challenge&code_challenge_method=plain&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=twitter` },
        { name: 'twitch', icon: TwitchLogo, status: false,  url: `${import.meta.env.VITE_TWITCH_CLIENT_API}?client_id=${import.meta.env.VITE_TWITCH_CLIENT_ID}&scope=user%3Aedit user%3Aread%3Aemail channel%3Amanage%3Apolls moderator%3Amanage%3Aannouncements&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=twitch` },
        { name: 'notion', icon: NotionLogo, status: false,  url: `${import.meta.env.VITE_NOTION_CLIENT_API}?client_id=${import.meta.env.VITE_NOTION_CLIENT_ID}&owner=user&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=notion` },
        { name: 'dropbox', icon: DropboxLogo, status: false,  url: `${import.meta.env.VITE_DROPBOX_CLIENT_API}?client_id=${import.meta.env.VITE_DROPBOX_CLIENT_ID}&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&response_type=code&state=dropbox` },
        { name: 'linkedin', icon: LinkedinLogo, status: false,  url: `${import.meta.env.VITE_LINKEDIN_CLIENT_API}?response_type=code&client_id=${import.meta.env.VITE_LINKEDIN_CLIENT_ID}&redirect_uri=${import.meta.env.VITE_SERVICES_REDIRECT_URI}&state=linkedin&scope=profile%20email%20openid%20w_member_social` },
    ];

    interface AvailableServices {
        name: string;
    }

    function updateServiceStatus(availableServices: AvailableServices[]) {
    const availableServiceMap = new Map(availableServices.map(service => [service.name, service]));
    services = services.map(service => ({
        ...service,
        status: availableServiceMap.has(service.name)
    }));
}


    onMount(async() => {
        pageTitle.set("Services");

        servicesStorage.subscribe(value => {
            if (value?.success) {
                console.log("testsuccess");
                gsap.to(`[data-service="${value.success}"]`, { delay: 0.5, backgroundColor: '#BAD2B7', duration: 1, onComplete: () => {
                    gsap.to(`[data-service="${value.success}"]`, { backgroundColor: 'rgba(255, 255, 255, 0.4)', duration: 1 });
                }});
            } else if (value?.failed) {
                gsap.to(`[data-service="${value.failed}"]`, { delay: 0.5, backgroundColor: '#EA8F99', duration: 1, onComplete: () => {
                    gsap.to(`[data-service="${value.failed}"]`, { backgroundColor: 'white', duration: 1 });
                }});
            }
        })

        const avaiblableServices = await getServices();
        updateServiceStatus(avaiblableServices);
    });

    afterUpdate(() => {
        ScrollTrigger.batch(".card", {
            interval: 0.1,
            batchMax: 3,
            onEnter: batch => gsap.to(batch, {autoAlpha: 1, stagger: 0.1, ease: "power2.inOut"}),
            onEnterBack: batch => gsap.to(batch, {autoAlpha: 1, stagger: 0.15, ease: "power2.inOut"}),
        });
        ScrollTrigger.refresh();
    });
</script>

<div class="container">
    {#each services as service}
        <a class="card" class:connected={service.status} data-service={service.name} href={service.url}>
            <svelte:component this={service.icon} style="margin: 2rem; font-size: 4rem; color: black;" />
            {#if service.status === true}
                <h1>{service.name} connected</h1>
            {:else}
                <h1>Connect with {service.name}</h1>
            {/if}
        </a>
    {/each}
</div>

<style>
    .container {
        padding: 2rem;
        display: flex;
        flex-wrap: wrap;
        gap: 2rem;
        justify-content: space-around;
    }
    .card {
        background-color: white;
        border: 4px solid var(--color4);
        display: flex;
        min-height: 145px;
        width: calc(100% / 3 - 2rem); /* Adjust the width to be one third of the container width, minus the gap */
        opacity: 0;
        padding: 1rem;
        border-radius: 15px;
        justify-content: space-evenly;
        align-items: center;
        cursor: pointer;
        transition: transform 0.3s ease-in-out;
        box-sizing: border-box; /* Include padding and border in the element's width and height */
    }

    .card.connected {
        background-color: rgba(255, 255, 255, 0.4);
    }

    :global(a) {
        text-decoration: none;
    }

    h1 {
        color: var(--color3);
        font-size: 2rem; /* Start with a smaller font size */
        word-wrap: break-word; /* Ensure the text wraps and does not overflow */
    }

    .card:hover {
        font-size: 1.3rem;
        transform: scale(1.05);

    }

    @media (max-width: 1200px) {
        .card {
            width: calc(100% / 2 - 2rem); /* 50% width on medium screens */
        }
    }

    @media (max-width: 768px) {
        .card {
            width: calc(100% - 2rem); /* 100% width on small screens */
        }

        h1 {
            font-size: 1.5rem; /* Smaller font size on small screens */
        }
    }
</style>
