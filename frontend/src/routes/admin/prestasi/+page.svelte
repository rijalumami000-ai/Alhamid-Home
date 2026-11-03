<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Achievement } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Trophy } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let list = $state<Achievement[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Achievement | null>(null);

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let achievementToDelete = $state<Achievement | null>(null);
  let isDeleting = $state(false);

  let form = $state({
    title: '',
    year: new Date().getFullYear(),
    level: 'Nasional',
    category: 'MQK',
    recipient_name: '',
    description: '',
    image_url: '',
  });

  async function loadData() {
    try {
      list = await api.admin.getAchievements();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      title: '',
      year: new Date().getFullYear(),
      level: 'Nasional',
      category: 'MQK',
      recipient_name: '',
      description: '',
      image_url: '',
    };
    isModalOpen = true;
  }

  function openEdit(a: Achievement) {
    editing = a;
    form = {
      title: a.title,
      year: a.year,
      level: a.level,
      category: a.category,
      recipient_name: a.recipient_name,
      description: a.description,
      image_url: a.image_url,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateAchievement(editing.id, form);
      } else {
        await api.admin.createAchievement(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan prestasi');
    }
  }

  function promptDelete(a: Achievement) {
    achievementToDelete = a;
    isDeleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!achievementToDelete) return;
    isDeleting = true;
    try {
      await api.admin.deleteAchievement(achievementToDelete.id);
      isDeleteModalOpen = false;
      achievementToDelete = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus prestasi');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Prestasi Santri | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Prestasi Santri & Lembaga
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Dokumentasikan kejuaraan MQK, Olimpiade Sains, Bahasa Arab/Inggris, dan Seni.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
    >
      <Plus class="w-4 h-4" /> Tambah Prestasi
    </button>
  </div>

  <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs">
        <thead class="bg-slate-50 border-b border-slate-200 text-slate-500 font-bold uppercase tracking-wider text-[10px]">
          <tr>
            <th class="p-4">Prestasi</th>
            <th class="p-4">Tahun</th>
            <th class="p-4">Tingkat</th>
            <th class="p-4">Peraih</th>
            <th class="p-4 text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-100">
          {#each list as item}
            <tr class="hover:bg-slate-50/70 transition">
              <td class="p-4 flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-amber-50 text-amber-700 flex items-center justify-center shrink-0 border border-amber-200">
                  <Trophy class="w-5 h-5" />
                </div>
                <div>
                  <div class="font-bold text-slate-900">{item.title}</div>
                  <div class="text-[11px] text-slate-500 line-clamp-1">{item.description}</div>
                </div>
              </td>
              <td class="p-4 font-mono font-bold text-slate-700">{item.year}</td>
              <td class="p-4">
                <span class="px-2 py-0.5 rounded-md text-[10px] font-bold bg-teal-50 text-teal-800 border border-teal-200">
                  {item.level}
                </span>
              </td>
              <td class="p-4 font-medium text-slate-700">{item.recipient_name || '-'}</td>
              <td class="p-4 text-right">
                <div class="inline-flex items-center gap-1">
                  <button
                    onclick={() => openEdit(item)}
                    class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
                    title="Edit Prestasi"
                  >
                    <Edit2 class="w-4 h-4" />
                  </button>
                  <button
                    onclick={() => promptDelete(item)}
                    class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
                    title="Hapus Prestasi"
                  >
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
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
          {editing ? 'Edit Prestasi' : 'Tambah Prestasi'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="p_title" class="text-xs font-bold text-slate-700 uppercase">Nama Prestasi *</label>
          <input
            id="p_title"
            type="text"
            required
            bind:value={form.title}
            placeholder="Juara 1 MQK..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div class="space-y-1">
            <label for="p_year" class="text-xs font-bold text-slate-700 uppercase">Tahun</label>
            <input
              id="p_year"
              type="number"
              bind:value={form.year}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
            />
          </div>
          <div class="space-y-1">
            <label for="p_level" class="text-xs font-bold text-slate-700 uppercase">Tingkat</label>
            <select
              id="p_level"
              bind:value={form.level}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
            >
              <option value="Kabupaten">Kabupaten</option>
              <option value="Provinsi">Provinsi</option>
              <option value="Nasional">Nasional</option>
              <option value="Internasional">Internasional</option>
            </select>
          </div>
        </div>

        <div class="space-y-1">
          <label for="p_recipient" class="text-xs font-bold text-slate-700 uppercase">Nama Peraih / Santri / Tim</label>
          <input
            id="p_recipient"
            type="text"
            bind:value={form.recipient_name}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <DriveImageInput
          id="p_img"
          label="Foto Sertifikat / Dokumentasi Prestasi"
          bind:value={form.image_url}
          placeholder="https://drive.google.com/... atau https://..."
        />

        <div class="space-y-1">
          <label for="p_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi / Keterangan</label>
          <textarea
            id="p_desc"
            rows="3"
            bind:value={form.description}
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
            Simpan Prestasi
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title="Hapus Prestasi Santri"
  itemName={achievementToDelete?.title}
  message="Rekam jejak prestasi ini akan dihapus secara permanen dari sistem."
  {isDeleting}
  onConfirm={confirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    achievementToDelete = null;
  }}
/>
