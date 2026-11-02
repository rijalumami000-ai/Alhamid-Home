<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Facility } from '$lib/types';
  import { Building2, Layers } from 'lucide-svelte';

  let list = $state<Facility[]>([]);
  let selectedCategory = $state('Semua');

  onMount(async () => {
    try {
      list = await api.getFacilities();
    } catch (e) {
      console.error(e);
    }
  });

  let categories = $derived(['Semua', ...new Set(list.map((f) => f.category))]);

  let filteredList = $derived(
    selectedCategory === 'Semua' ? list : list.filter((f) => f.category === selectedCategory)
  );
</script>

<svelte:head>
  <title>Fasilitas Pesantren | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Sarana & Prasarana
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Fasilitas Kampus Pesantren
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Lingkungan pembelajaran asri, nyaman, dan memadai untuk menunjang kegiatan ibadah, studi akademik, hingga kebugaran santri.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16 space-y-8">
  <!-- Filter Categories -->
  <div class="flex flex-wrap items-center gap-2">
    {#each categories as cat}
      <button
        onclick={() => (selectedCategory = cat)}
        class="px-4 py-2 rounded-xl text-xs font-semibold transition {selectedCategory === cat
          ? 'bg-emerald-800 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        {cat}
      </button>
    {/each}
  </div>

  <!-- Facilities Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each filteredList as item}
      <div class="bg-white rounded-2xl border border-slate-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden group flex flex-col">
        <div class="relative h-56 overflow-hidden bg-slate-900">
          <img
            src={item.image_url}
            alt={item.name}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute top-3 left-3">
            <span class="px-2.5 py-1 rounded-full bg-slate-950/70 text-emerald-300 text-[11px] font-semibold backdrop-blur-md">
              {item.category}
            </span>
          </div>
        </div>
        <div class="p-6 flex-1 flex flex-col justify-between space-y-2">
          <h2 class="font-serif-luxury font-bold text-xl text-slate-900 group-hover:text-emerald-700 transition">
            {item.name}
          </h2>
          <p class="text-slate-600 text-xs sm:text-sm leading-relaxed">
            {item.description}
          </p>
        </div>
      </div>
    {/each}
  </div>
</div>

