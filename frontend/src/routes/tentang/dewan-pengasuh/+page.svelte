<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Caregiver } from '$lib/types';
  import { Shield, BookOpen, HeartHandshake } from 'lucide-svelte';

  let list = $state<Caregiver[]>([]);
  let loading = $state(true);

  onMount(async () => {
    try {
      list = await api.getCaregivers();
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Dewan Pengasuh & Masyayikh | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Masyayikh & Pendidik
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Dewan Pengasuh Pesantren
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Para ulama, masyayikh, dan asatidz yang mengasuh, mendidik, dan mengawal keaslian sanad keilmuan para santri Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-100 shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden flex flex-col sm:flex-row gap-6 p-6">
        <div class="w-full sm:w-44 h-56 sm:h-auto rounded-2xl overflow-hidden shrink-0 bg-slate-100">
          <img
            src={item.photo_url}
            alt={item.name}
            class="w-full h-full object-cover"
          />
        </div>
        <div class="flex-1 flex flex-col justify-between space-y-3">
          <div>
            <span class="text-xs font-bold text-emerald-700 uppercase tracking-wide">
              {item.role_title}
            </span>
            <h2 class="font-serif-luxury font-bold text-xl text-slate-900 mt-1">
              {item.name}
            </h2>
            <p class="text-slate-600 text-xs sm:text-sm mt-3 leading-relaxed">
              {item.bio}
            </p>
          </div>

          {#if item.sanad_keilmuan}
            <div class="pt-3 border-t border-slate-100">
              <div class="text-[11px] font-bold text-amber-700 uppercase tracking-wider flex items-center gap-1 mb-1">
                <BookOpen class="w-3 h-3" /> Sanad Keilmuan
              </div>
              <p class="text-xs text-slate-500 italic">
                "{item.sanad_keilmuan}"
              </p>
            </div>
          {/if}
        </div>
      </div>
    {/each}
  </div>
</div>

