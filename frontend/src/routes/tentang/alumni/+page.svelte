<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Alumni } from '$lib/types';
  import {
    MapPin,
    GraduationCap,
    Search,
    X,
    ChevronLeft,
    ChevronRight,
    Maximize2,
    Calendar,
    Sparkles
  } from 'lucide-svelte';

  let list = $state<Alumni[]>([]);
  let loading = $state(true);
  let selectedYear = $state<string>('all');
  let searchQuery = $state<string>('');
  let selectedAlumni = $state<Alumni | null>(null);

  onMount(async () => {
    try {
      list = await api.getAlumni();
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  });

  // Extract unique sorted graduation years
  let availableYears = $derived.by(() => {
    const years = Array.from(new Set(list.map((item) => item.graduation_year))).filter(Boolean);
    return years.sort((a, b) => b - a);
  });

  // Filtered alumni based on year and search
  let filteredList = $derived.by(() => {
    let result = list;
    if (selectedYear !== 'all') {
      const yr = parseInt(selectedYear, 10);
      result = result.filter((item) => item.graduation_year === yr);
    }
    if (searchQuery.trim()) {
      const q = searchQuery.toLowerCase().trim();
      result = result.filter((item) => {
        const nameMatch = item.name.toLowerCase().includes(q);
        const locMatch = (item.current_location || '').toLowerCase().includes(q);
        const unitMatch = (item.education_unit || '').toLowerCase().includes(q);
        return nameMatch || locMatch || unitMatch;
      });
    }
    return result;
  });

  // Group alumni by graduation year
  let groupedAlumni = $derived.by(() => {
    const groups: { year: number; items: Alumni[] }[] = [];
    const targetYears = selectedYear === 'all'
      ? availableYears
      : [parseInt(selectedYear, 10)];

    for (const year of targetYears) {
      const items = filteredList.filter((item) => item.graduation_year === year);
      if (items.length > 0) {
        groups.push({ year, items });
      }
    }
    return groups;
  });

  // Lightbox navigation helpers
  function openLightbox(alumni: Alumni) {
    selectedAlumni = alumni;
  }

  function closeLightbox() {
    selectedAlumni = null;
  }

  function nextAlumni() {
    if (!selectedAlumni) return;
    const currentIndex = filteredList.findIndex((a) => a.id === selectedAlumni!.id);
    if (currentIndex !== -1 && currentIndex < filteredList.length - 1) {
      selectedAlumni = filteredList[currentIndex + 1];
    } else if (filteredList.length > 0) {
      selectedAlumni = filteredList[0];
    }
  }

  function prevAlumni() {
    if (!selectedAlumni) return;
    const currentIndex = filteredList.findIndex((a) => a.id === selectedAlumni!.id);
    if (currentIndex > 0) {
      selectedAlumni = filteredList[currentIndex - 1];
    } else if (filteredList.length > 0) {
      selectedAlumni = filteredList[filteredList.length - 1];
    }
  }

  function handleKeydown(e: KeyboardEvent) {
    if (!selectedAlumni) return;
    if (e.key === 'Escape') closeLightbox();
    if (e.key === 'ArrowRight') nextAlumni();
    if (e.key === 'ArrowLeft') prevAlumni();
  }
</script>

<svelte:window onkeydown={handleKeydown} />

<svelte:head>
  <title>Ikatan Alumni Al-Hamid (IKAH) | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header Banner -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-teal-400">
      Jejaring Santri & Alumni
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Ikatan Alumni Pondok Pesantren Al-Hamid
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Direktori santri lulusan lintas angkatan yang mengabdi dan berkontribusi bagi masyarakat, membawa sanad keilmuan dan akhlak luhur pesantren.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16 space-y-12">
  <!-- Interactive Angkatan Filter Bar & Search -->
  <div class="bg-white rounded-3xl border border-slate-200/80 p-6 sm:p-8 shadow-sm space-y-6">
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
      <div>
        <h2 class="font-display font-bold text-xl text-slate-900">
          Pilih Angkatan Kelulusan
        </h2>
        <p class="text-xs text-slate-500 mt-0.5">
          Filter alumni berdasarkan tahun angkatan atau cari berdasarkan nama santri dan daerah asal.
        </p>
      </div>

      <!-- Quick Search Input -->
      <div class="relative w-full md:w-72">
        <Search class="w-4 h-4 text-slate-400 absolute left-3.5 top-1/2 -translate-y-1/2" />
        <input
          type="text"
          placeholder="Cari nama atau daerah..."
          bind:value={searchQuery}
          class="w-full pl-10 pr-4 py-2.5 rounded-2xl bg-slate-50 border border-slate-200 text-xs font-medium text-slate-900 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-teal-600 focus:bg-white transition"
        />
        {#if searchQuery}
          <button
            onclick={() => (searchQuery = '')}
            class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600"
          >
            <X class="w-3.5 h-3.5" />
          </button>
        {/if}
      </div>
    </div>

    <!-- Angkatan Pill Selector Tabs -->
    <div class="flex flex-wrap items-center gap-2 pt-2 border-t border-slate-100">
      <button
        onclick={() => (selectedYear = 'all')}
        class="px-4 py-2 rounded-xl text-xs font-mono font-semibold transition-all duration-200 flex items-center gap-2 {selectedYear === 'all' ? 'bg-slate-950 text-white shadow-md' : 'bg-slate-100 text-slate-600 hover:bg-slate-200'}"
      >
        <span>Semua Angkatan</span>
        <span class="text-[10px] opacity-75 font-normal">({list.length})</span>
      </button>

      {#each availableYears as year}
        {@const count = list.filter((i) => i.graduation_year === year).length}
        <button
          onclick={() => (selectedYear = String(year))}
          class="px-4 py-2 rounded-xl text-xs font-mono font-semibold transition-all duration-200 flex items-center gap-2 {selectedYear === String(year) ? 'bg-teal-700 text-white shadow-md shadow-teal-900/20' : 'bg-slate-100 text-slate-700 hover:bg-slate-200'}"
        >
          <span>Angkatan {year}</span>
          <span class="text-[10px] opacity-75 font-normal">({count})</span>
        </button>
      {/each}
    </div>
  </div>

  <!-- Alumni Grouped Display -->
  {#if loading}
    <div class="text-center py-20 text-slate-400 text-sm font-mono">
      Memuat direktori alumni pesantren...
    </div>
  {:else if groupedAlumni.length === 0}
    <div class="text-center py-20 bg-white rounded-3xl border border-slate-200 p-8 space-y-3">
      <div class="font-display font-bold text-lg text-slate-800">Tidak ada data alumni ditemukan</div>
      <p class="text-xs text-slate-500 max-w-sm mx-auto">
        Silakan coba kata kunci lain atau pilih opsi 'Semua Angkatan'.
      </p>
    </div>
  {:else}
    <div class="space-y-16">
      {#each groupedAlumni as group}
        <section class="space-y-8" id="angkatan-{group.year}">
          <!-- Distinct Angkatan Separator & Header -->
          <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 pb-4 border-b-2 border-slate-200/90">
            <div class="space-y-1">
              <div class="flex items-center gap-3">
                <h2 class="font-display font-black text-2xl sm:text-3xl text-slate-950">
                  Angkatan {group.year}
                </h2>
                <span class="text-xs font-mono font-semibold text-teal-700 bg-teal-50 px-2.5 py-1 rounded-md border border-teal-200/60">
                  Lulusan Tahun {group.year}
                </span>
              </div>
              <p class="text-xs text-slate-500">
                Pondok Pesantren Al-Hamid &bull; MTs / MA Cintamulya
              </p>
            </div>

            <div class="text-xs font-mono text-slate-500 font-medium">
              {group.items.length} Santri Terdata
            </div>
          </div>

          <!-- Alumni Cards Grid for this Angkatan -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 sm:gap-8">
            {#each group.items as item}
              <div
                class="bg-white rounded-3xl border border-slate-200/80 p-6 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 flex flex-col justify-between space-y-5 group"
              >
                <div class="space-y-4">
                  <!-- Photo with Zoom Indicator -->
                  <button
                    type="button"
                    class="relative overflow-hidden rounded-2xl bg-slate-950 aspect-[4/3] w-full block text-left group/btn"
                    onclick={() => openLightbox(item)}
                    aria-label="Lihat foto penuh {item.name}"
                  >
                    <img
                      src={item.photo_url || 'https://images.unsplash.com/photo-1544717305-2782549b5136?auto=format&fit=crop&w=600&q=80'}
                      alt={item.name}
                      class="w-full h-full object-cover group-hover/btn:scale-105 transition duration-500"
                    />
                    <div class="absolute inset-0 bg-gradient-to-t from-slate-950/70 via-transparent to-transparent opacity-0 group-hover/btn:opacity-100 transition duration-300 flex items-end justify-between p-4">
                      <span class="text-xs font-mono text-white flex items-center gap-1.5 drop-shadow">
                        <Maximize2 class="w-3.5 h-3.5 text-teal-300" /> Lihat Foto Penuh
                      </span>
                    </div>
                  </button>

                  <!-- Alumni Details (No Profession / Career!) -->
                  <div class="space-y-1.5">
                    <button
                      type="button"
                      class="text-left w-full block"
                      onclick={() => openLightbox(item)}
                    >
                      <h3 class="font-display font-bold text-lg text-slate-950 hover:text-teal-700 transition leading-snug">
                        {item.name}
                      </h3>
                    </button>
                    <div class="flex items-center gap-2 text-xs text-slate-500 font-medium">
                      <GraduationCap class="w-3.5 h-3.5 text-teal-600 shrink-0" />
                      <span>{item.education_unit || 'Alumni Al-Hamid'}</span>
                    </div>
                    {#if item.current_location}
                      <div class="flex items-center gap-2 text-xs text-slate-500 font-medium">
                        <MapPin class="w-3.5 h-3.5 text-teal-600 shrink-0" />
                        <span>Asal: {item.current_location}</span>
                      </div>
                    {/if}
                  </div>

                  <!-- Santri Quote / Reflection -->
                  {#if item.story}
                    <p class="font-editorial italic text-xs text-slate-600 pt-3 border-t border-slate-100 leading-relaxed line-clamp-3">
                      "{item.story}"
                    </p>
                  {/if}
                </div>

                <!-- Footer Action to enlarge -->
                <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
                  <span class="text-[11px] font-mono text-slate-400">Angkatan {item.graduation_year}</span>
                  <button
                    onclick={() => openLightbox(item)}
                    class="text-xs font-display font-bold text-teal-700 hover:text-teal-900 inline-flex items-center gap-1 transition"
                  >
                    <span>Perbesar Foto</span>
                    <Maximize2 class="w-3 h-3" />
                  </button>
                </div>
              </div>
            {/each}
          </div>
        </section>
      {/each}
    </div>
  {/if}
</div>

<!-- ================= HIGH-RESOLUTION LIGHTBOX MODAL ================= -->
{#if selectedAlumni}
  <div
    class="fixed inset-0 z-50 bg-slate-950/90 backdrop-blur-md flex items-center justify-center p-4 sm:p-8 animate-in fade-in duration-200"
    role="dialog"
    aria-modal="true"
    tabindex="-1"
    onkeydown={(e) => {
      if (e.key === 'Escape') closeLightbox();
    }}
    onclick={(e) => {
      if (e.target === e.currentTarget) closeLightbox();
    }}
  >
    <!-- Modal Card Container -->
    <div class="bg-[#0B101B] border border-slate-800 rounded-3xl max-w-4xl w-full overflow-hidden shadow-2xl relative flex flex-col lg:flex-row text-white max-h-[90vh]">
      <!-- Close Button -->
      <button
        onclick={closeLightbox}
        class="absolute top-4 right-4 z-20 w-10 h-10 rounded-full bg-slate-900/80 hover:bg-slate-800 text-slate-300 hover:text-white flex items-center justify-center transition border border-slate-700"
        aria-label="Tutup"
      >
        <X class="w-5 h-5" />
      </button>

      <!-- Navigation Arrows -->
      <button
        onclick={prevAlumni}
        class="absolute left-4 top-1/2 -translate-y-1/2 z-20 w-10 h-10 rounded-full bg-slate-900/80 hover:bg-slate-800 text-slate-300 hover:text-white flex items-center justify-center transition border border-slate-700"
        aria-label="Foto Sebelumnya"
      >
        <ChevronLeft class="w-5 h-5" />
      </button>

      <button
        onclick={nextAlumni}
        class="absolute right-4 top-1/2 -translate-y-1/2 z-20 w-10 h-10 rounded-full bg-slate-900/80 hover:bg-slate-800 text-slate-300 hover:text-white flex items-center justify-center transition border border-slate-700"
        aria-label="Foto Berikutnya"
      >
        <ChevronRight class="w-5 h-5" />
      </button>

      <!-- Large Image Display -->
      <div class="lg:w-3/5 bg-black flex items-center justify-center overflow-hidden min-h-[350px] lg:min-h-[500px]">
        <img
          src={selectedAlumni.photo_url || 'https://images.unsplash.com/photo-1544717305-2782549b5136?auto=format&fit=crop&w=1200&q=90'}
          alt={selectedAlumni.name}
          class="w-full h-full object-contain max-h-[75vh]"
        />
      </div>

      <!-- Detail Story Panel -->
      <div class="lg:w-2/5 p-8 flex flex-col justify-between space-y-6 overflow-y-auto">
        <div class="space-y-4">
          <div class="flex items-center gap-2">
            <span class="text-xs font-mono font-semibold px-2.5 py-1 rounded bg-teal-500/20 text-teal-300 border border-teal-500/30">
              Angkatan {selectedAlumni.graduation_year}
            </span>
            <span class="text-xs text-slate-400 font-mono">
              Tahun Kelulusan {selectedAlumni.graduation_year}
            </span>
          </div>

          <h3 class="font-display font-black text-2xl sm:text-3xl text-white leading-tight">
            {selectedAlumni.name}
          </h3>

          <div class="space-y-2 text-xs text-slate-300 pt-2 border-t border-slate-800">
            <div class="flex items-center gap-2">
              <GraduationCap class="w-4 h-4 text-teal-400 shrink-0" />
              <span>{selectedAlumni.education_unit || 'Pondok Pesantren Al-Hamid'}</span>
            </div>
            {#if selectedAlumni.current_location}
              <div class="flex items-center gap-2">
                <MapPin class="w-4 h-4 text-teal-400 shrink-0" />
                <span>Asal Daerah: {selectedAlumni.current_location}</span>
              </div>
            {/if}
          </div>

          {#if selectedAlumni.story}
            <div class="pt-4 border-t border-slate-800">
              <div class="text-[11px] font-mono uppercase tracking-wider text-teal-400 font-semibold mb-2">
                Pesan & Refleksi Santri
              </div>
              <p class="font-editorial italic text-sm text-slate-300 leading-relaxed">
                "{selectedAlumni.story}"
              </p>
            </div>
          {/if}
        </div>

        <div class="pt-4 border-t border-slate-800 text-[11px] text-slate-500 flex items-center justify-between font-mono">
          <span>Gunakan tombol panah keyboard ← →</span>
          <span>ESC untuk tutup</span>
        </div>
      </div>
    </div>
  </div>
{/if}
