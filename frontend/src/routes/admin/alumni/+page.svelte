<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Alumni } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Users } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let list = $state<Alumni[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Alumni | null>(null);

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let alumniToDelete = $state<Alumni | null>(null);
  let isDeleting = $state(false);

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

  function promptDelete(item: Alumni) {
    alumniToDelete = item;
    isDeleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!alumniToDelete) return;
    isDeleting = true;
    try {
      await api.admin.deleteAlumni(alumniToDelete.id);
      isDeleteModalOpen = false;
      alumniToDelete = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus alumni');
    } finally {
      isDeleting = false;
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
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Ikatan Alumni Al-Hamid (IKAPHA)
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola basis data alumni per angkatan dan rekam jejak lulusan.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
    >
      <Plus class="w-4 h-4" /> Tambah Alumni
    </button>
  </div>

  <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs">
        <thead class="bg-slate-50 border-b border-slate-200 text-slate-500 font-bold uppercase tracking-wider text-[10px]">
          <tr>
            <th class="p-4">Alumni</th>
            <th class="p-4">Angkatan (Tahun Lulus)</th>
            <th class="p-4">Unit Pendidikan</th>
            <th class="p-4">Domisili</th>
            <th class="p-4 text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-100">
          {#each list as item}
            <tr class="hover:bg-slate-50/70 transition">
              <td class="p-4 flex items-center gap-3">
                <div class="w-10 h-10 rounded-full bg-slate-200 overflow-hidden shrink-0 border border-slate-200">
                  {#if item.photo_url}
                    <img src={item.photo_url} alt={item.name} class="w-full h-full object-cover" />
                  {:else}
                    <div class="w-full h-full flex items-center justify-center text-slate-400">
                      <Users class="w-4 h-4" />
                    </div>
                  {/if}
                </div>
                <div>
                  <div class="font-bold text-slate-900">{item.name}</div>
                  {#if item.story}
                    <div class="text-[11px] text-slate-500 line-clamp-1">{item.story}</div>
                  {/if}
                </div>
              </td>
              <td class="p-4 font-mono font-bold text-teal-700">Angkatan {item.graduation_year}</td>
              <td class="p-4 font-medium text-slate-700">{item.education_unit || '-'}</td>
              <td class="p-4 text-slate-600">{item.current_location || '-'}</td>
              <td class="p-4 text-right">
                <button
                  onclick={() => openEdit(item)}
                  class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition mr-1"
                  title="Edit Alumni"
                >
                  <Edit2 class="w-4 h-4" />
                </button>
                <button
                  onclick={() => promptDelete(item)}
                  class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
                  title="Hapus Alumni"
                >
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
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editing ? 'Edit Alumni' : 'Tambah Alumni'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
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
            <label for="al_year" class="text-xs font-bold text-slate-700 uppercase">Tahun Kelulusan *</label>
            <input
              id="al_year"
              type="number"
              required
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

        <DriveImageInput
          id="al_photo"
          label="Foto Alumni"
          bind:value={form.photo_url}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="al_loc" class="text-xs font-bold text-slate-700 uppercase">Asal Daerah / Kota Domisili</label>
          <input
            id="al_loc"
            type="text"
            bind:value={form.current_location}
            placeholder="Lampung Selatan / Bandar Lampung / Jakarta"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="al_story" class="text-xs font-bold text-slate-700 uppercase">Pesan & Kesan Alumni</label>
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
            class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Alumni
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title="Hapus Alumni"
  itemName={alumniToDelete?.name}
  message="Data alumni ini akan dihapus dari arsip direktori alumni pesantren."
  {isDeleting}
  onConfirm={confirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    alumniToDelete = null;
  }}
/>
