<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Facility } from '$lib/types';
  import { Building2, Layers, Sparkles } from 'lucide-svelte';

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

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-teal-400">
      Infrastruktur & Sarana
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Fasilitas Kampus Pesantren
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Lingkungan pembelajaran asri, representatif, dan memadai untuk menunjang kegiatan ibadah, studi sains, teknologi, serta kebugaran santri.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20 space-y-10">
  <!-- Filter Categories -->
  <div class="flex flex-wrap items-center gap-2">
    {#each categories as cat}
      <button
        onclick={() => (selectedCategory = cat)}
        class="px-5 py-2 rounded-full text-xs font-display font-bold transition {selectedCategory === cat
          ? 'bg-slate-900 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        {cat}
      </button>
    {/each}
  </div>

  <!-- Facilities Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each filteredList as item}
      <div class="bg-white rounded-3xl border border-slate-200/80 shadow-md hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 overflow-hidden flex flex-col group">
        <div class="relative h-56 overflow-hidden bg-slate-950">
          <img
            src={item.image_url}
            alt={item.name}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-slate-950/70 via-transparent to-transparent"></div>
          <div class="absolute top-3 left-3">
            <span class="px-3 py-1 rounded-full bg-slate-900/90 text-teal-300 text-[11px] font-bold backdrop-blur-md border border-white/10">
              {item.category}
            </span>
          </div>
        </div>
        <div class="p-6 flex-1 flex flex-col justify-between space-y-3">
          <div>
            <h2 class="font-display font-black text-xl text-slate-950 group-hover:text-teal-700 transition leading-snug">
              {item.name}
            </h2>
            <p class="text-slate-600 text-xs sm:text-sm mt-2 leading-relaxed">
              {item.description}
            </p>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>
