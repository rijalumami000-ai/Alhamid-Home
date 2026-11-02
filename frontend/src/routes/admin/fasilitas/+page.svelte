<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Facility } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Building2 } from 'lucide-svelte';

  let list = $state<Facility[]>([]);
  let isModalOpen = $state(false);
  let editing = $state<Facility | null>(null);
  let form = $state({
    name: '',
    category: 'Akademik',
    description: '',
    image_url: '',
    order: 0,
  });

  async function loadData() {
    try {
      list = await api.admin.getFacilities();
    } catch (e) {
      console.error(e);
    }
  }

  function openCreate() {
    editing = null;
    form = {
      name: '',
      category: 'Akademik',
      description: '',
      image_url: 'https://images.unsplash.com/photo-1581092918056-0c4c3acd3789?auto=format&fit=crop&w=800&q=80',
      order: list.length + 1,
    };
    isModalOpen = true;
  }

  function openEdit(f: Facility) {
    editing = f;
    form = {
      name: f.name,
      category: f.category,
      description: f.description,
      image_url: f.image_url,
      order: f.order,
    };
    isModalOpen = true;
  }

  async function handleSave(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editing) {
        await api.admin.updateFacility(editing.id, form);
      } else {
        await api.admin.createFacility(form);
      }
      isModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan fasilitas');
    }
  }

  async function handleDelete(id: number) {
    if (confirm('Hapus fasilitas ini?')) {
      try {
        await api.admin.deleteFacility(id);
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
  <title>Kelola Fasilitas | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
        Fasilitas & Sarana Pesantren
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Kelola gedung, laboratorium, asrama, sarana ibadah, dan sarana olahraga.
      </p>
    </div>
    <button
      onclick={openCreate}
      class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-bold shadow-md transition"
    >
      <Plus class="w-4 h-4" /> Tambah Fasilitas
    </button>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each list as item}
      <div class="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col justify-between">
        <div>
          <div class="h-44 bg-slate-900 relative">
            <img src={item.image_url} alt={item.name} class="w-full h-full object-cover" />
            <span class="absolute top-3 left-3 px-2 py-0.5 rounded-full bg-slate-900/80 text-white text-[10px] font-bold uppercase">
              {item.category}
            </span>
          </div>
          <div class="p-5 space-y-2">
            <h2 class="font-bold text-base text-slate-900">{item.name}</h2>
            <p class="text-xs text-slate-500 line-clamp-3">{item.description}</p>
          </div>
        </div>

        <div class="p-4 border-t border-slate-100 flex items-center justify-between bg-slate-50/50">
          <span class="text-xs text-slate-400">Urutan: #{item.order}</span>
          <div class="space-x-1">
            <button
              onclick={() => openEdit(item)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-emerald-700 hover:bg-emerald-50"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => handleDelete(item.id)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-red-600 hover:bg-red-50"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

<!-- MODAL FORM -->
{#if isModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-serif-luxury font-bold text-lg text-slate-900">
          {editing ? 'Edit Fasilitas' : 'Tambah Fasilitas'}
        </h2>
        <button onclick={() => (isModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSave} class="p-6 space-y-4">
        <div class="space-y-1">
          <label for="f_name" class="text-xs font-bold text-slate-700 uppercase">Nama Fasilitas *</label>
          <input
            id="f_name"
            type="text"
            required
            bind:value={form.name}
            placeholder="Contoh: Lab Komputer MAI"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="f_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori</label>
          <select
            id="f_cat"
            bind:value={form.category}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium bg-white"
          >
            <option value="Ibadah">Ibadah</option>
            <option value="Asrama">Asrama</option>
            <option value="Akademik">Akademik</option>
            <option value="Olahraga">Olahraga</option>
            <option value="Penunjang">Penunjang</option>
          </select>
        </div>

        <div class="space-y-1">
          <label for="f_img" class="text-xs font-bold text-slate-700 uppercase">URL Gambar</label>
          <input
            id="f_img"
            type="url"
            bind:value={form.image_url}
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>

        <div class="space-y-1">
          <label for="f_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi Fasilitas</label>
          <textarea
            id="f_desc"
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
            class="px-5 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold shadow-md"
          >
            Simpan Fasilitas
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

