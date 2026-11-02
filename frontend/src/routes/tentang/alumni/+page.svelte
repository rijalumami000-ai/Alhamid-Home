<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Alumni } from '$lib/types';
  import { Users, MapPin, Briefcase, GraduationCap, ArrowUpRight } from 'lucide-svelte';

  let list = $state<Alumni[]>([]);

  onMount(async () => {
    try {
      list = await api.getAlumni();
    } catch (e) {
      console.error(e);
    }
  });
</script>

<svelte:head>
  <title>Ikatan Alumni Al-Hamid (IKAH) | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Jejaring & Pengabdian
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Ikatan Alumni Pondok Pesantren Al-Hamid
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Wadah sinergi ribuan alumni Al-Hamid yang berkiprah di sektor keagamaan, akademisi, sains teknologi, birokrasi, dan wirausaha di berbagai belahan dunia.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-100 p-6 shadow-sm hover:shadow-xl transition-all duration-300 flex flex-col justify-between space-y-4">
        <div class="space-y-4">
          <div class="flex items-center gap-4">
            <img
              src={item.photo_url}
              alt={item.name}
              class="w-16 h-16 rounded-2xl object-cover border-2 border-emerald-600 shadow-sm"
            />
            <div>
              <h2 class="font-bold text-base text-slate-900">{item.name}</h2>
              <div class="text-xs text-emerald-700 font-semibold">Angkatan {item.graduation_year}</div>
              <div class="text-[11px] text-slate-400">{item.education_unit}</div>
            </div>
          </div>

          <div class="space-y-1.5 text-xs text-slate-600">
            <div class="flex items-center gap-2">
              <Briefcase class="w-3.5 h-3.5 text-slate-400" />
              <span>{item.profession}</span>
            </div>
            <div class="flex items-center gap-2">
              <MapPin class="w-3.5 h-3.5 text-slate-400" />
              <span>{item.current_location}</span>
            </div>
          </div>

          <p class="text-slate-600 text-xs leading-relaxed italic bg-slate-50 p-3 rounded-xl border border-slate-100">
            "{item.story}"
          </p>
        </div>

        {#if item.social_link}
          <div class="pt-3 border-t border-slate-100 flex justify-end">
            <a
              href={item.social_link}
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center gap-1 text-xs font-semibold text-emerald-700 hover:text-emerald-800"
            >
              <span>Profil LinkedIn / Medsos</span>
              <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        {/if}
      </div>
    {/each}
  </div>
</div>

