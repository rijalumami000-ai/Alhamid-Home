<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Alumni } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Users } from 'lucide-svelte';

  let list = $state<Alumni[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Alumni | null>(null);
  let form = $state({
    name: '',
    graduation_year: 2020,
    education_unit: 'MA Cintamulya',
    profession: '',
    current_location: '',
    story: '',
    photo_url: '',
    social_link: '',
    order: 0,
  });

  async function loadData() {
    try {
      list = await api.admin.getAlumni();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      name: '',
      graduation_year: 2022,
      education_unit: 'MA Cintamulya',
      profession: '',
      current_location: '',
      story: '',
      photo_url: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=500&q=80',
      social_link: '',
      order: list.length + 1,
    };
    isModalOpen = true;
  }

  function openEdit(item: Alumni) {
    editing = item;
    form = {
      name: item.name,
      graduation_year: item.graduation_year,
      education_unit: item.education_unit,
      profession: item.profession,
      current_location: item.current_location,
      story: item.story,
      photo_url: item.photo_url,
      social_link: item.social_link,
      order: item.order,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateAlumni(editing.id, form);
      } else {
        await api.admin.createAlumni(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan data alumni');
    }
  }

  async function handleDelete(id: number) {
    if (confirm('Hapus profil alumni ini?')) {
      try {
        await api.admin.deleteAlumni(id);
        await loadData();
      } catch (err: any) {
        alert(err.message || 'Gagal menghapus');
      }
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Ikatan Alumni | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Direktori Ikatan Alumni (IKAH)
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola profil kiprah lulusan, profesi, dan lokasi alumni Al-Hamid.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
    >
      <Plus class="w-4 h-4" /> Tambah Alumni
    </button>
  </div>

  <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs">
        <thead class="bg-slate-50 border-b border-slate-200 text-slate-500 font-bold uppercase tracking-wider">
          <tr>
            <th class="p-4">Nama Alumni</th>
            <th class="p-4">Angkatan</th>
            <th class="p-4">Unit Lulusan</th>
            <th class="p-4">Profesi & Domisili</th>
            <th class="p-4 text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-100">
          {#each list as item}
            <tr class="hover:bg-slate-50/70 transition">
              <td class="p-4">
                <div class="flex items-center gap-3">
                  <img src={item.photo_url} alt={item.name} class="w-9 h-9 rounded-xl object-cover border border-slate-200" />
                  <div>
                    <div class="font-bold text-slate-900">{item.name}</div>
                    <div class="text-[11px] text-slate-400 line-clamp-1">{item.story}</div>
                  </div>
                </div>
              </td>
              <td class="p-4 font-mono font-bold text-slate-700">{item.graduation_year}</td>
              <td class="p-4 text-slate-600 font-medium">{item.education_unit}</td>
              <td class="p-4 text-slate-600">
                <div class="font-medium text-slate-800">{item.profession}</div>
                <div class="text-[11px] text-slate-400">{item.current_location}</div>
              </td>
              <td class="p-4 text-right space-x-1">
                <button onclick={() => openEdit(item)} class="p-1.5 rounded-lg text-slate-500 hover:text-emerald-700">
                  <Edit2 class="w-4 h-4" />
                </button>
                <button onclick={() => handleDelete(item.id)} class="p-1.5 rounded-lg text-slate-500 hover:text-red-600">
                  <Trash2 class="w-4 h-4" />
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">
          {editing ? 'Edit Alumni' : 'Tambah Alumni'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4">
        <div class="space-y-1">
          <label for="al_name" class="text-xs font-bold text-slate-700 uppercase">Nama Lengkap *</label>
          <input
            id="al_name"
            type="text"
            required
            bind:value={form.name}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="al_year" class="text-xs font-bold text-slate-700 uppercase">Tahun Kelulusan</label>
            <input
              id="al_year"
              type="number"
              bind:value={form.graduation_year}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="al_unit" class="text-xs font-bold text-slate-700 uppercase">Unit Pendidikan</label>
            <input
              id="al_unit"
              type="text"
              bind:value={form.education_unit}
              placeholder="MA Cintamulya / Diniyah"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="al_prof" class="text-xs font-bold text-slate-700 uppercase">Profesi / Karier</label>
            <input
              id="al_prof"
              type="text"
              bind:value={form.profession}
              placeholder="Dokter / Dosen / IT"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="al_loc" class="text-xs font-bold text-slate-700 uppercase">Domisili Sekarang</label>
            <input
              id="al_loc"
              type="text"
              bind:value={form.current_location}
              placeholder="Bandung / Kairo"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
        </div>

        <div class="space-y-1">
          <label for="al_story" class="text-xs font-bold text-slate-700 uppercase">Cerita Kiprah / Kesan</label>
          <textarea
            id="al_story"
            rows="3"
            bind:value={form.story}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>

        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button
            type="button"
            onclick={() => (isModalOpen = false)}
            class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md"
          >
            Simpan Alumni
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

