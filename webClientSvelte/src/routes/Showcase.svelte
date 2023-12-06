<script lang="ts">
    import gsap from 'gsap';
    import { ScrollTrigger } from 'gsap/ScrollTrigger';
    import { onMount } from 'svelte';
    import MouseScrollIcon from 'virtual:icons/iconoir/mouse-scroll-wheel';
    import ShowcaseNavbar from "../components/common/ShowcaseNavbar.svelte";
    gsap.registerPlugin(ScrollTrigger);

    onMount(() => {
        const texts = Array.from(document.querySelectorAll<HTMLElement>('.text, .fixed-text'));
        let throttleTimeout: any = null;

        const updateOpacity = () => {
            const viewportHeight = window.innerHeight;
            const offset = viewportHeight * 0.1;
            const fadeBoundary = viewportHeight * 0.3;

            texts.forEach((text: HTMLElement) => {
                const rect = text.getBoundingClientRect();
                const distanceToCenter = Math.abs((viewportHeight / 2 - offset) - rect.top);
                const opacityValue = Math.max(0, Math.min(1, (fadeBoundary - distanceToCenter) / fadeBoundary));
                gsap.to(text, { opacity: opacityValue });
            });
        };

        gsap.timeline({ delay: 1, repeat: -1, repeatDelay: 2 })
        .to(".mouse-scroll", {
            scale: 1.5,
            yPercent: -50,
            duration: 1,
            ease: "power3.inOut"
        })
        .to(".mouse-scroll", {
            scale: 1,
            duration: 1,
            yPercent: 0,
            ease: "power3.inOut"
        })

        gsap.to(".text, .fixed-text", {
            yPercent: -3000,
            scale: 2.5,
            stagger: { amount: 1 },
            scrollTrigger: {
                trigger: ".starwars-container",
                scrub: 2,
                pin: true,
                end: "+=300%",
                onUpdate: () => {
                    if (throttleTimeout) {
                        clearTimeout(throttleTimeout);
                    }
                    throttleTimeout = setTimeout(updateOpacity, 10);
                },
                onScrubComplete: () => {
                    if (throttleTimeout) {
                        clearTimeout(throttleTimeout);
                    }
                    throttleTimeout = setTimeout(updateOpacity, 10);
                },
            }
        });
    });
</script>

<div class='showcase-container'>
    <ShowcaseNavbar />
    <div class='starwars-container'>
            <p class="fixed-text">Build custom workflows in minutes</p>
            <p class="fixed-text"><br/></p>
            <p class="fixed-text"><br/></p>
            <p class="text">Automate the busywork</p>
            <p class="text">so you can focus on your job, not your tools</p>
            <p class="text">We'll show you how</p>
            <p class="text"><br/></p>
            <p class="text"><br/></p>
            <p class="text">Get more power from your tools</p>
            <p class="text"><br/></p>
            <p class="text"><br/></p>
            <p class="text">Integrate your critical work apps into workflows</p>
            <p class="text">reclaim your time, and focus on impactful work</p>
            <p class="text"><br/></p>
            <p class="text"><br/></p>
            <p class="text">Connect the apps you already love</p>
            <p class="text"><br/></p>
            <p class="text"><br/></p>
            <p class="text">Area supports more apps than any other platform</p>
            <p class="text">so you can optimize the tools you use</p>
    </div>
    <button class="mouse-scroll">
        <MouseScrollIcon />
    </button>
</div>

<style>
    .showcase-container {
        position: relative;
        width: 100%;
        height: 100%;
        background-color: var(--backgroundColor);
    }

    .showcase-container .starwars-container {
        position: relative;
        text-align: center;
        height: 100vh;
        overflow: hidden;
    }

    .showcase-container .fixed-text {
        position: relative;
        color: white;
        font-size: 3rem;
        top: 40%;
        white-space: nowrap;
    }
    .showcase-container .text {
        position: absolute;
        color: white;
        font-size: 1.5rem;
        left: 50%;
        top: 110%;
        transform: translateX(-50%) translateY(-50%);
        white-space: nowrap;
    }

    .showcase-container .mouse-scroll {
        position: fixed;
        font-size: 2rem;
        top: 90%;
        left: 50%;
        transform: translateX(-50%);
        color: white;
        background: transparent;
        border: none;
        cursor: pointer;
        transition: color 0.3s ease-in-out;
    }

    .showcase-container .mouse-scroll:hover {
        color: var(--color2);
    }

    @media (max-width: 1200px) {
        .starwars-container {
            font-size: 1rem;
        }
    }
</style>
