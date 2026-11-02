<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Testimonial } from '$lib/types';
  import { Star, MessageSquareQuote } from 'lucide-svelte';

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

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Amanah & Kepercayaan
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Testimoni Alumni & Wali Santri
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Cerita dan refleksi nyata mengenai pengalaman belajar, pembentukan adab, dan bekal hidup yang didapatkan di Pondok Pesantren Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-100 p-6 sm:p-8 shadow-sm hover:shadow-xl transition-all duration-300 flex flex-col justify-between space-y-6">
        <div class="space-y-4">
          <div class="flex items-center gap-1 text-amber-400">
            {#each Array(item.rating || 5) as _}
              <Star class="w-4 h-4 fill-amber-400" />
            {/each}
          </div>
          <p class="text-slate-600 text-sm leading-relaxed italic">
            "{item.content}"
          </p>
        </div>

        <div class="pt-4 border-t border-slate-100 flex items-center gap-3">
          <img
            src={item.avatar_url}
            alt={item.name}
            class="w-12 h-12 rounded-full object-cover border-2 border-emerald-700 shadow-sm"
          />
          <div>
            <h2 class="font-bold text-sm text-slate-800">{item.name}</h2>
            <div class="text-xs text-slate-500">{item.role}</div>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

