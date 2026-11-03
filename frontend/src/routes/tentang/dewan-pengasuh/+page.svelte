<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Caregiver } from '$lib/types';
  import { Shield, BookOpen, HeartHandshake, Sparkles } from 'lucide-svelte';

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

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
      <Sparkles class="w-3.5 h-3.5 text-teal-400" />
      <span>Masyayikh & Dewan Pembina</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Dewan Pengasuh Pesantren
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Para ulama, masyayikh, dan asatidz yang mengasuh, membimbing akhlak, dan mengawal keaslian transmisi sanad keilmuan para santri Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200/80 shadow-md hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 overflow-hidden flex flex-col sm:flex-row gap-6 p-6 group">
        <div class="w-full sm:w-48 h-60 sm:h-auto rounded-2xl overflow-hidden shrink-0 bg-slate-950">
          <img
            src={item.photo_url}
            alt={item.name}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
        </div>
        <div class="flex-1 flex flex-col justify-between space-y-4">
          <div>
            <span class="text-[11px] font-black text-teal-700 uppercase tracking-widest">
              {item.role_title}
            </span>
            <h2 class="font-display font-black text-xl text-slate-950 mt-1 leading-snug group-hover:text-teal-700 transition">
              {item.name}
            </h2>
            <p class="text-slate-600 text-xs sm:text-sm mt-3 leading-relaxed">
              {item.bio}
            </p>
          </div>

          {#if item.sanad_keilmuan}
            <div class="pt-3 border-t border-slate-100 flex items-start gap-2 text-xs text-slate-500">
              <Shield class="w-4 h-4 text-teal-600 shrink-0 mt-0.5" />
              <span class="font-medium text-[11px] leading-snug">{item.sanad_keilmuan}</span>
            </div>
          {/if}
        </div>
      </div>
    {/each}
  </div>
</div>
