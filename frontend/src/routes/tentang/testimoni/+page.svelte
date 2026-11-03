<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Testimonial } from '$lib/types';
  import { Star, MessageSquareQuote, Sparkles } from 'lucide-svelte';

  let list = $state<Testimonial[]>([]);

  onMount(async () => {
    try {
      list = await api.getTestimonials();
    } catch (e) {
      console.error(e);
    }
  });
</script>

<svelte:head>
  <title>Testimoni Alumni & Wali Santri | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
      <Sparkles class="w-3.5 h-3.5 text-teal-400" />
      <span>Amanah & Kepercayaan</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Testimoni Alumni & Wali Santri
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Cerita dan refleksi nyata mengenai pengalaman belajar, pembentukan adab, dan bekal hidup yang didapatkan di Pondok Pesantren Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200/80 p-8 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 flex flex-col justify-between space-y-6">
        <div class="space-y-4">
          <div class="flex items-center gap-1 text-amber-400">
            {#each Array(item.rating || 5) as _}
              <Star class="w-4 h-4 fill-amber-400" />
            {/each}
          </div>
          <p class="font-editorial italic text-slate-700 text-base leading-relaxed">
            "{item.content}"
          </p>
        </div>

        <div class="pt-4 border-t border-slate-100 flex items-center gap-3.5">
          <img
            src={item.avatar_url}
            alt={item.name}
            class="w-12 h-12 rounded-full object-cover border-2 border-teal-500 shadow-sm"
          />
          <div>
            <h2 class="font-display font-bold text-sm text-slate-950">{item.name}</h2>
            <div class="text-xs text-slate-500">{item.role}</div>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>
